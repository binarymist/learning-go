package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revError := Reverse(input)
	doubleRev, doubleRevError := Reverse(rev)
	fmt.Printf("original: %v\n", input)
	fmt.Printf("reversed: %v, error: %v\n", rev, revError)
	fmt.Printf("reversed again: %v, error: %v\n", doubleRev, doubleRevError)
}
