package stringutil

import "unicode"

// ToUpper uppercases all the runes in its argument string.
func ToUpper(s string) string {
	r := []rune(s)
	for index, value := range r {
		r[index] = unicode.ToUpper(value)
	}
	return string(r)
}
