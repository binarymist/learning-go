// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Reverb1 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"unicode"
)

//!+
func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	if len(shout) != 0 {
		fmt.Fprintln(conn, "\t", string(unicode.ToUpper(rune(shout[0])))+shout[1:])
	} else {
		fmt.Fprintln(conn, "\t", "")
	}
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() { // Wait for messages from the client.
		// Docs for net.Conn https://pkg.go.dev/net#Conn state:
		// Multiple goroutines may invoke methods on a Conn simultaneously.
		go echo(conn, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	conn.Close()
}

//!-

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
