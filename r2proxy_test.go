package main

import (
	"reflect"
	"testing"
)

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
