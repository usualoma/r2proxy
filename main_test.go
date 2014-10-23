package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestCmdOptions(t *testing.T) {
	opts := cmdOptions{
		OptAllowedPorts:    []int{80, 443},
		OptDestinationHost: "www.example.com",
		OptDestinationPort: 81,
		OptVerbose:         true,
	}

	if !reflect.DeepEqual(opts.AllowedPorts(), []int{80, 443}) {
		t.Errorf("got %q", opts.AllowedPorts())
	}

	if opts.DestinationHost() != "www.example.com" {
		t.Errorf("got %q", opts.DestinationHost())
	}

	if opts.DestinationPort() != 81 {
		t.Errorf("got %d", opts.DestinationPort())
	}

	if opts.Verbose() != true {
		t.Errorf("got %q", opts.Verbose())
	}
}

func TestRunInvalidArgs(t *testing.T) {
	var b bytes.Buffer
	Run([]string{"arg"}, &b)
	if strings.Index(b.String(), "Usage") == -1 {
		t.Errorf("got %q", b.String())
	}
}

func TestRunHelp(t *testing.T) {
	var b bytes.Buffer
	Run([]string{"--help"}, &b)
	if strings.Index(b.String(), "Usage") == -1 {
		t.Errorf("got %q", b.String())
	}
}

func TestRunVersion(t *testing.T) {
	var b bytes.Buffer
	Run([]string{"--version"}, &b)
	if strings.Index(b.String(), "r2proxy:") != 0 {
		t.Errorf("got %q", b.String())
	}
}
