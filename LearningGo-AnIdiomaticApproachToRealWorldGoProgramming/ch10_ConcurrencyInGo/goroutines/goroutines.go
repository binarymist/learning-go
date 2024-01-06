package main

import (
	"fmt"
)

func process(val int) int {
	// do something with val
	return val // Placeholder processing logic (double the input value).
}

func insertIntValuesToChannel(vals []int, out chan<- int) {
	go func() {
		for _, val := range vals {
			out <- val
		}
		close(out)
	}()
}

func runThingConcurrently(out chan<- int, in <-chan int) {
	go func() {
		for val := range in {
			result := process(val)
			out <- result
		}
		close(out)
	}()
}

func main() {
	// Create channels
	input := make(chan int)
	output := make(chan int)

	vals := []int{1, 2, 3, 4, 5}
	insertIntValuesToChannel(vals, input)
	// Start the concurrent processing goroutine
	runThingConcurrently(output, input)

	// Receive and print results
	for result := range output {
		fmt.Println(result)
	}
}
