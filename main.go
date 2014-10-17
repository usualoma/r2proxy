package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"os"
)

var version = "v0.1.0"

type cmdOptions struct {
	OptHelp            bool   `short:"h" long:"help" description:"Show this help message and exit"`
	OptVersion         bool   `long:"version" description:"Print the version and exit"`
	OptVerbose         bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	OptListenPort      int    `long:"listen-port" default:"8080" description:"Listen port"`
	OptAllowedPorts    []int  `long:"allowed-dest-port" default:"80" default:"443" description:"Destination port(s) that will be allowed"`
	OptDestinationHost string `long:"fixed-dest-host" description:"Fixed destination host"`
	OptDestinationPort int    `long:"fixed-dest-port" description:"Fixed destination port"`
}

func (o cmdOptions) AllowedPorts() []int {
	return o.OptAllowedPorts
}

func (o cmdOptions) DestinationHost() string {
	return o.OptDestinationHost
}

func (o cmdOptions) DestinationPort() int {
	return o.OptDestinationPort
}

func (o cmdOptions) Verbose() bool {
	return o.OptVerbose
}

func main() {
	var err error
	var st int

	defer func() { os.Exit(st) }()

	opts := &cmdOptions{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.Parse()
	if len(args) > 0 || err != nil {
		st = 1
	}

	if st != 0 || opts.OptHelp {
		p.WriteHelp(os.Stderr)
		return
	}

	if opts.OptVersion {
		fmt.Fprintf(os.Stderr, "r2proxy: %s\n", version)
		return
	}

	proxy := NewProxyHttpServer(opts)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", opts.OptListenPort), proxy))
}
