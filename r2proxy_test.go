package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

type TestProxyOptions struct {
	OptAllowedPorts    []int
	OptDestinationHost string
	OptDestinationPort int
	OptVerbose         bool
}

func (o TestProxyOptions) AllowedPorts() []int {
	return o.OptAllowedPorts
}

func (o TestProxyOptions) DestinationHost() string {
	return o.OptDestinationHost
}

func (o TestProxyOptions) DestinationPort() int {
	return o.OptDestinationPort
}

func (o TestProxyOptions) Verbose() bool {
	return o.OptVerbose
}

type TestResponseWriter struct {
	buffer *[]byte
}

func (w TestResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w TestResponseWriter) Write(bytes []byte) (int, error) {
	*w.buffer = append(*w.buffer, bytes...)
	return len(bytes), nil
}

func (w TestResponseWriter) WriteHeader(i int) {
}

func init() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Path[1:])
	})
	go (func() {
		http.ListenAndServe(":8980", handler)
	})()
}

func TestNormalizePorts(t *testing.T) {
	actual := NormalizePorts([]int{1, 3, 5})
	expected := map[int]bool{
		1: true,
		3: true,
		5: true,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %q\nwant %q", actual, expected)
	}
}

func TestNormalizeHost(t *testing.T) {
	suites := map[string]string{
		"127.0.0.1":             "127.0.0.1",
		"www.exsample.com":      "www.exsample.com",
		"www.exsample.com:8080": "www.exsample.com",
	}

	for host, expected := range suites {
		actual := NormalizeHost(host)
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

func NewWriterReqest() (TestResponseWriter, *http.Request) {
	b := new([]byte)
	w := TestResponseWriter{buffer: b}
	r := new(http.Request)
	r.RemoteAddr = "127.0.0.1:9999"
	r.Header = http.Header{}

	return w, r
}

func TestNewProxyHttpServerGET(t *testing.T) {
	expected := "OK"

	opts := TestProxyOptions{
		OptAllowedPorts:    []int{8980, 443},
		OptDestinationHost: "",
		OptDestinationPort: 0,
		OptVerbose:         false,
	}

	w, r := NewWriterReqest()

	r.URL, _ = url.Parse("http://www.example.com:8980/" + expected)

	server := NewProxyHttpServer(opts)
	server.ServeHTTP(w, r)

	if !bytes.Equal(*w.buffer, []byte(expected)) {
		t.Errorf("got %q\nwant %q", *w.buffer, expected)
	}
}

func TestNewProxyHttpServerDenied(t *testing.T) {
	expected := "You are not allowed to connect to this port: 80"

	opts := TestProxyOptions{
		OptAllowedPorts:    []int{443},
		OptDestinationHost: "",
		OptDestinationPort: 0,
		OptVerbose:         false,
	}

	w, r := NewWriterReqest()

	r.URL, _ = url.Parse("http://www.example.com/OK")

	server := NewProxyHttpServer(opts)
	server.ServeHTTP(w, r)

	if !bytes.Equal(*w.buffer, []byte(expected)) {
		t.Errorf("got %q\nwant %q", *w.buffer, expected)
	}
}

func TestNewProxyHttpServerFixedHost(t *testing.T) {
	expected := "OK"

	opts := TestProxyOptions{
		OptAllowedPorts:    []int{8980, 443},
		OptDestinationHost: "localhost:50",
		OptDestinationPort: 0,
		OptVerbose:         false,
	}

	w, r := NewWriterReqest()

	r.URL, _ = url.Parse("http://www.example.com:8980/" + expected)

	server := NewProxyHttpServer(opts)
	server.ServeHTTP(w, r)

	if !bytes.Equal(*w.buffer, []byte(expected)) {
		t.Errorf("got %q\nwant %q", *w.buffer, expected)
	}
}
