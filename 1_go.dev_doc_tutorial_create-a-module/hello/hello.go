package main

import (
  "fmt"
  "log"
  "go.dev_doc_tutorial/greetings"
)

// https://go.dev/doc/tutorial/call-module-code
// > go mod edit -replace example.com/greetings=../greetings
// > go mod tidy
// > go run .

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// names := []string{"Maximus", "Kim", "Leanne"}
	var names []string
  names = append(names, "Maximus", "Kim", "Leanne")
  messages, err := greetings.Hellos(names)
	if err != nil { log.Fatal(err) }
  fmt.Println(messages)
}