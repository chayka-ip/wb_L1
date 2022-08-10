package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.

func main() {
	s := "snow dog sun"
	fmt.Println(s)
	fmt.Println(ReverseWords(s))
}

func ReverseWords(s string) string {
	/*
		Source string will be split into words
		based on spaces and special characters.

		Result string contains words written in reverse order
		Separated by whitespaces
	*/
	w := strings.Fields(s)
	out := ""

	for i := (len(w) - 1); i >= 0; i-- {
		out += w[i]
		if i > 0 {
			out += " "
		}
	}
	return out
}
