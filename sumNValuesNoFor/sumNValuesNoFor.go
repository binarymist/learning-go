// Challenge:
//   Calculate the sum of squares of n numbers without for loops, and only using the standard library.
//   https://softwareengineering.stackexchange.com/questions/314150/flow-control-in-go-without-a-for-loop

//   One goroutine should call another goroutine to do the simple squaring.
//   The first goroutine should keep track of the number of times the second goroutine has been called.

package main

import (
	"fmt"
)

// Recursive
func square(array []int, index int) []int {
	if index == len(array) {
		return array
	}
	array[index] = array[index] * array[index]
	return square(array, index+1)
}

// Recursive
func sumArray(array []int, index int) int {
	if index == len(array)-1 {
		return array[index]
	}
	return array[index] + sumArray(array, index+1)
}

func sumSquares(numsToSqr []int, sumOfNumsSquared chan int) {
	// Let's not mutate:
	deepCopyOfNumsToSqr := append([]int{}, numsToSqr...)
	elementsSquared := square(deepCopyOfNumsToSqr, 0)
	sum := sumArray(elementsSquared, 0)
	sumOfNumsSquared <- sum
}

func main() {
	// myArray := [5]int{} // Fixed size
	// mySlice := []int{} // Dynamically resized
	numbersToSquare := []int{1, 2, 3, 4, 5}
	sumOfNumbersSquared := make(chan int)
	go sumSquares(numbersToSquare, sumOfNumbersSquared)
	result := <-sumOfNumbersSquared
	close(sumOfNumbersSquared)
	fmt.Println("Sum of squares of", numbersToSquare, "values is", result)
}
