package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var receivedTimes = make(map[string]chan string)

func handleServer(serverAddr string, location string) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Printf("Failed to connect to %s: %s\n", location, err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Printf("Error reading from %s: %s\n", location, err)
			}
			return
		}
		receivedTime := strings.TrimSpace(string(buffer[:bytesRead]))

		receivedTimes[location] <- receivedTime
	}
}

func printReceivedTime(countries []string) {
	clearScreen := "\033[2J" // ANSI escape sequence to clear screen

	for {
		fmt.Print(clearScreen) // Clear the screen
		fmt.Print("\033[H")    // Move cursor to home position

		fmt.Println("Received Times:")
		fmt.Println("---------------")
		fmt.Printf("| %-10s | %-25s |\n", "Location", "Time")
		fmt.Println("---------------")

		for _, country := range countries {
			time := <-receivedTimes[country]
			fmt.Printf("| %-10s | %-25s |\n", country, time)
		}

		time.Sleep(time.Second)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: clockwall country1=host1:port1 country2=host1:port1 ...")
		return
	}

	countries := []string{}

	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			fmt.Printf("Invalid argument format: %s\n", arg)
			return
		}
		country := parts[0]
		address := parts[1]
		countries = append(countries, country)
		receivedTimes[country] = make(chan string)
		go handleServer(address, country)
	}

	go printReceivedTime(countries)

	// To keep the program running
	select {}
}
