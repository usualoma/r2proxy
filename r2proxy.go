package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"github.com/elazarl/goproxy"
)

type ProxyOptions interface {
	AllowedPorts() []int
	DestinationHost() string
	DestinationPort() int
	Verbose() bool
}

func NormalizePorts(ports []int) map[int]bool {
    var m = map[int]bool{}
    for _, a := range ports { m[a] = true }
    return m
}

func NormalizeHost(address string) string {
	if strings.Contains(address, ":") {
		return strings.Split(address, ":")[0]
	} else {
        return address
    }
}

func NewProxyHttpServer(opts ProxyOptions) http.HandlerFunc {
	portRegexp := regexp.MustCompile(":([0-9]+)$")

	allowedPortMap := NormalizePorts(opts.AllowedPorts())
	destinationHost := NormalizeHost(opts.DestinationHost())

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = opts.Verbose()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := destinationHost
		if host == "" {
			host = strings.Split(r.RemoteAddr, ":")[0]
		}

		port := opts.DestinationPort()
		if port == 0 {
			matches := portRegexp.FindStringSubmatch(r.URL.Host)
			if len(matches) != 0 {
				port, _ = strconv.Atoi(matches[1])
			} else {
				port = 80
			}

			if ! allowedPortMap[port] {
				fmt.Fprintf(w, "You are not allowed to connect to this port: %d", port)
				return
			}
		}
		r.URL.Host = fmt.Sprintf("%s:%d", host, port)

		proxy.ServeHTTP(w, r)
	})
}
