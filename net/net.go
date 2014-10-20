package net

import (
	"io"
	"net"
	"os"
	"strconv"

	"github.com/felixrabe/limited-netcat/errors"
)

// Listen waits for another host to connect.
func Listen(host string, port int) {
	l, err := net.Listen("tcp", hostPortToString(host, port))
	if err != nil {
		errors.Fatal(err)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		errors.Fatal(err)
	}
	defer conn.Close()

	if err := entangle(conn); err != nil {
		errors.Fatal(err)
	}
}

// Connect connects to another host.
func Connect(host string, port int) {
	conn, err := net.Dial("tcp", hostPortToString(host, port))
	if err != nil {
		errors.Fatal(err)
	}
	defer conn.Close()

	if err := entangle(conn); err != nil {
		errors.Fatal(err)
	}
}

func hostPortToString(host string, port int) string {
	return host + ":" + strconv.Itoa(port)
}

// entangle connects the io.ReadWriter to standard input and output.
func entangle(rw io.ReadWriter) error {
	errc := make(chan error, 1)
	go cp(os.Stdout, rw, errc)
	go cp(rw, os.Stdin, errc)
	return <-errc
}

// cp returns the result of io.Copy over the channel.
func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}
