package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
	Разработать программу, которая переворачивает подаваемую на ход строку
	Символы могут быть unicode.
*/

func main() {
	str := "Hello, 世界界"
	fmt.Println(str)
	fmt.Println(ReverseString4(str))
}

func ReverseString(s string) string {
	runes := []rune{}
	b := []byte(s)

	/*
		utf8 package is used to read runes from string represented as byte slice.
		Rune can be encoded with 1, 2, 3 or 4 bytes.
		Package is able to determine number of runes and their size
	*/

	// while there are some undecoded runes
	for utf8.RuneCount(b) > 0 {
		// decode rune from the end
		r, size := utf8.DecodeLastRune(b)
		// shrink byte slice
		b = b[:len(b)-size]
		// update result
		runes = append(runes, r)
	}
	return string(runes)
}

func ReverseString2(s string) string {
	// split string into individual characters
	d := strings.Split(s, "")
	out := ""
	// iterate over string slice from the end to start
	for i := len(d) - 1; i >= 0; i-- {
		//update result
		out += d[i]
	}
	return out
}

func ReverseString3(s string) string {
	out := []rune{}
	src := []rune(s)
	for i := len(src) - 1; i >= 0; i-- {
		out = append(out, src[i])
	}
	return string(out)
}

func ReverseString4(s string) string {
	b := strings.Builder{}
	src := []rune(s)
	for i := len(src) - 1; i >= 0; i-- {
		b.WriteRune(src[i])
	}
	return b.String()
}
