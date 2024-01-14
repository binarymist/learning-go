// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// When we get responses from the network (conn), io.Copy reads from src and writes to dst.
	go mustCopy(os.Stdout, conn)
	// Read from standard input and send to the network (conn).
	mustCopy(conn, os.Stdin)
	// Docs for net.Conn https://pkg.go.dev/net#Conn state:
	// Multiple goroutines may invoke methods on a Conn simultaneously.
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
