package main

import (
	"strings"
)

var justString string

// creates unicode string
func createHugeString(n int) string {
	return strings.Repeat("也", n)
}

func someFunc() {
	v := createHugeString(1 << 10)

	/*
		This is wrong approach to get slice of characters from  unicode sting:

		justString = v[:100]
		Output: 也也也也也也也也也也也也也也也也也也也也也也也也也也也也也也也也也�


		GO string is a sequence of runes that can be presented as byte slice.
		It is known that rune have size of 1, 2, 3 or 4 bytes.

		Attempt to slice string directly when there are runes
		with size more than 1 byte leads to:
		 1) data corruption;
		 2) wrong number of actual characters in the output;
	*/

	// convert source string to rune slice
	rs := []rune(v)
	// get sequence of runes with desired length and convert it back to string
	justString = string(rs[:100])
}
func main() {
	someFunc()
}
