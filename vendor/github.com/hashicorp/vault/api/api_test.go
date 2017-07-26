package api

import (
	"fmt"
	"net"
	"net/http"
	"testing"

	"golang.org/x/net/http2"
)

// testHTTPServer creates a test HTTP server that handles requests until
// the listener returned is closed.
func testHTTPServer(
	t *testing.T, handler http.Handler) (*Config, net.Listener) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	server := &http.Server{Handler: handler}
	if err := http2.ConfigureServer(server, nil); err != nil {
		t.Fatal(err)
	}
	go server.Serve(ln)

	config := DefaultConfig()
	config.Address = fmt.Sprintf("http://%s", ln.Addr())

	return config, ln
}
