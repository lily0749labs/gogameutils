package httpUtil

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestJSONWithClient(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodPost {
			t.Errorf("request method = %q, want POST", request.Method)
		}
		if got := request.Header.Get("Content-Type"); got != "application/json;charset=UTF-8" {
			t.Errorf("Content-Type = %q", got)
		}
		if got := request.Header.Get("X-Game"); got != "hall" {
			t.Errorf("X-Game = %q, want hall", got)
		}
		body, err := io.ReadAll(request.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}
		if got := string(body); got != `{"phone":"911234567890"}` {
			t.Errorf("request body = %q", got)
		}

		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader(`{"accepted":false}`)),
			Header:     make(http.Header),
		}, nil
	})}

	body, err := JSONWithClient(client, "post", "https://example.invalid", map[string]string{
		"phone": "911234567890",
	}, map[string]string{"X-Game": "hall"})
	if err != nil {
		t.Fatalf("JSONWithClient() error = %v", err)
	}
	if got := string(body); got != `{"accepted":false}` {
		t.Fatalf("JSONWithClient() body = %q", got)
	}
}

func TestJSONWithClientMarshalError(t *testing.T) {
	_, err := JSONWithClient(http.DefaultClient, http.MethodPost, "https://example.invalid", make(chan int), nil)
	if err == nil {
		t.Fatal("JSONWithClient() error = nil, want JSON marshal error")
	}
}
