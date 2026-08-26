package tcp

import (
	"net"
	"testing"

	"github.com/lily0749labs/gogameutils/tcp/inter"
)

type testServer struct{}

var _ inter.Server = (*testServer)(nil)

func (*testServer) Accept(net.Conn) {}
func (*testServer) Close(net.Conn)  {}
func (*testServer) Error(error)     {}
func (*testServer) Exit()           {}

func TestNewTcpClient(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen() error = %v", err)
	}
	defer listener.Close()

	accepted := make(chan error, 1)
	go func() {
		conn, err := listener.Accept()
		if err == nil {
			conn.Close()
		}
		accepted <- err
	}()

	conn, err := NewTcpClient(listener.Addr().String())
	if err != nil {
		t.Fatalf("NewTcpClient() error = %v", err)
	}
	conn.Close()
	if err := <-accepted; err != nil {
		t.Fatalf("Accept() error = %v", err)
	}
}

func TestNewTcpServer(t *testing.T) {
	server := NewTcpServer("127.0.0.1", 8080, &testServer{})
	if server.ip != "127.0.0.1" || server.port != 8080 {
		t.Fatalf("NewTcpServer() address = %s:%d", server.ip, server.port)
	}
}

func TestNewTcpServerRejectsInvalidAddress(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		port int
	}{
		{name: "invalid IP", ip: "not-an-ip", port: 8080},
		{name: "invalid port", ip: "127.0.0.1", port: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("NewTcpServer() did not panic")
				}
			}()
			NewTcpServer(tc.ip, tc.port, &testServer{})
		})
	}
}
