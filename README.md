# r2proxy [![Build Status](https://travis-ci.org/usualoma/r2proxy.svg?branch=master)](https://travis-ci.org/usualoma/r2proxy) [![Coverage Status](https://img.shields.io/coveralls/usualoma/r2proxy.svg)](https://coveralls.io/r/usualoma/r2proxy?branch=master)

The r2proxy (this name means "reflective reverse proxy") is an implementation of the proxy server. This program sends all the request to direct-remote-host.


## Use case

![Use case](https://raw.githubusercontent.com/usualoma/r2proxy/master/artwork/use-case.png)

* We can use the "proxy.pac" for a HTTP request, in order to access an ELB instance for a host that has not been associated to DNS record.
    * A HTTP listener of the ELB instance can handle proxy request.
* We cannot use the "proxy.pac" for a HTTPS request easily.
    * A HTTP listener of the ELB instance can not handle CONNECT method.
* The r2proxy works as a proxy server to direct-remote-host, and can handle COONECT method.
    * An ELB instance responds even for a HTTP listener port of internal interface. (Oct. 2014)
* We can use the "proxy.pac" for a HTTPS request via the TCP listener and the r2proxy!


## Features

* Can use a ELB's HTTPS listener.
    * A configuration of HTTPS is unnecessary in an EC2 instance.
* Can use the ELB's "Cookie Stickiness".
* A configuration is unnecessary in the case of the standard usage.


## Installation

### Binary

Binary packages are available in the [releases page](https://github.com/usualoma/r2proxy/releases).

### go get

```
go get github.com/usualoma/r2proxy
```


## Setup

1. Add a TCP listener to an ELB instance.
1. Run r2proxy at an EC2 instance.
1. Add proxy settings of HTTPS to "proxy.pac".


## Command-Line Options

```
Usage:
  r2proxy

Application Options:
  -h, --help               Show this help message and exit
      --version            Print the version and exit
  -v, --verbose            Show verbose debug information
      --listen-port=       Listen port (8080)
      --allowed-dest-port= Destination port(s) that will be allowed (80, 443)
      --fixed-dest-host=   Fixed destination host
      --fixed-dest-port=   Fixed destination port
```

## LICENSE

Copyright (c) 2014 Taku AMANO

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
