// Limited netcat in Go.
//
// Inspired by http://vimeo.com/53221560.
//
// Installation:
//
//     go get github.com/felixrabe/limited-netcat
//
// Usage:
//
//     limited-netcat -l 0.0.0.0 1234
//     limited-netcat localhost 1234
package main

import (
	"github.com/felixrabe/limited-netcat/flags"
	"github.com/felixrabe/limited-netcat/net"
)

func main() {
	listen, host, port := flags.Parse()
	if listen {
		net.Listen(host, port)
	} else {
		net.Connect(host, port)
	}
}
