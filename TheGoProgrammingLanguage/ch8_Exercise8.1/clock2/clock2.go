package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		timestamp := time.Now().Format("15:04:05")
		_, err := fmt.Fprintf(c, "\r%s", timestamp)
		// _, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // writes the time to the client
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var port int
	portZeroValue := 0

	flag.IntVar(&port, "port", portZeroValue, "Provide a port number")
	flag.IntVar(&port, "p", portZeroValue, "Provide a port number (short)")
	flag.Parse()

	if port == 0 {
		fmt.Println("Please provide a value for the port flag")
		flag.Usage()
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port)) // Creates a net.Listener
	if err != nil {
		if port < 1024 {
			log.Printf("Port: %d is a privileged port. Elevated permissions might be required.", port)
		}
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port:", listener.Addr().(*net.TCPAddr).Port)
	for {
		// block until an incoming connection request is received
		// , then returns a net.Conn object representing the connection
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
