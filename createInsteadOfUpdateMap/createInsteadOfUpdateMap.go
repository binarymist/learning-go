// Code from https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4
//   This is about creating rather than updating/mutating maps

// After considerable investigation, Go has a spread operator
// but it only appears to be able to be used on arrays.
// Lets try convert the maps to arrays and use the spread to create a new array .......
//   Even doing this requires a for loop, in-fact every solution I've seen requires for loops. WTF?
//   Does this language not have any functional tools?
package main

import (
	"fmt"
	"rsc.io/quote"
)

func mergeMaps(mapA, mapB map[string]int) map[string]int {
	allAges := make(map[string]int, len(mapA)+len(mapB))
	for k, v := range mapA {
		allAges[k] = v
	}
	for k, v := range mapB {
		allAges[k] = v
	}
	return allAges
}

func main() {
	ages1 := map[string]int{"John": 30}
	ages2 := map[string]int{"Mary": 28}

	allAges := mergeMaps(ages1, ages2)

	fmt.Println("allAges is: ", allAges)

	fmt.Println(quote.Go())
}
