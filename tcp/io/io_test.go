package io

import (
	"net"
	"reflect"
	"testing"
	"time"
)

func TestWriteRead(t *testing.T) {
	server, client := net.Pipe()
	t.Cleanup(func() {
		server.Close()
		client.Close()
	})

	written := make(chan struct {
		n   int
		err error
	}, 1)
	go func() {
		n, err := Write(client, []byte("hello"))
		written <- struct {
			n   int
			err error
		}{n: n, err: err}
	}()

	n, data, err := Read(server)
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if n != len(data) || string(data) != "hello" {
		t.Fatalf("Read() = (%d, %q), want (5, %q)", n, data, "hello")
	}

	result := <-written
	if result.err != nil {
		t.Fatalf("Write() error = %v", result.err)
	}
	if result.n != 4+len(data) {
		t.Fatalf("Write() bytes = %d, want %d", result.n, 4+len(data))
	}
}

type testPayload struct {
	Number int
	Text   string
}

func TestJSONRoundTrip(t *testing.T) {
	want := testPayload{Number: 7, Text: "hello"}
	got := testPayload{}
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	errCh := make(chan error, 1)
	go func() {
		errCh <- WriteJsonByTime(client, want, time.Now().Add(time.Second))
	}()
	if err := ReadJsonByTime(server, &got, time.Now().Add(time.Second)); err != nil {
		t.Fatalf("ReadJsonByTime() error = %v", err)
	}
	if err := <-errCh; err != nil {
		t.Fatalf("WriteJsonByTime() error = %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("JSON round trip = %#v, want %#v", got, want)
	}
}

func TestGobRoundTrip(t *testing.T) {
	want := testPayload{Number: 9, Text: "world"}
	got := testPayload{}
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	errCh := make(chan error, 1)
	go func() {
		errCh <- WriteGobByTime(client, want, time.Now().Add(time.Second))
	}()
	if err := ReadGobByTime(server, &got, time.Now().Add(time.Second)); err != nil {
		t.Fatalf("ReadGobByTime() error = %v", err)
	}
	if err := <-errCh; err != nil {
		t.Fatalf("WriteGobByTime() error = %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Gob round trip = %#v, want %#v", got, want)
	}
}

func TestIntegerRoundTrip(t *testing.T) {
	for _, want := range []int{0, 1, -1, 1<<31 - 1, -1 << 31} {
		if got := BytesToInt(IntToBytes(want)); got != want {
			t.Fatalf("BytesToInt(IntToBytes(%d)) = %d", want, got)
		}
	}
}
