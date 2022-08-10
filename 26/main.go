package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	Разработать программу, которая проверяет, что все символы в строке
	уникальные (true — если уникальные, false etc). Функция проверки должна быть
	регистронезависимой.
*/

func main() {
	s := []string{"abcd", "abCdefAaf", "aabcd", "a A", "aAvs"}
	for _, v := range s {
		fmt.Printf("- %s [is unique: %t]\n", v, IsUniqueString1(v))
	}

	/*
			Output:
		- abcd [is unique: true]
		- abCdefAaf [is unique: false]
		- aabcd [is unique: false]
		- a A [is unique: false]
		- aAvs [is unique: false]
	*/
}

// case is ignored
func IsUniqueString1(s string) bool {
	/*
		Map is used to store runes found in string.
		If duplicate is found - string is not unique.
	*/
	m := make(map[string]bool)
	/*
		String is converted to lower (because case should be ignored)
		and split to the slice of individual characters

	*/
	p := strings.Split(strings.ToLower(s), "")
	for _, v := range p {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}

// case is ignored
func IsUniqueString2(s string) bool {
	// String is converted to lower case because case should be ignored.
	s = strings.ToLower(s)
	p := strings.Split(s, "")
	for _, v := range p {
		// If character presents more than once in the string - string is not unique.
		if strings.Count(s, v) > 1 {
			return false
		}
	}
	return true
}

// case is ignored
func IsUniqueString3(s string) bool {
	/*
		Map is used to store runes found in string.
		If duplicate is found - string is not unique.
	*/
	m := make(map[rune]bool)
	// iterate over each rune in source string
	for _, v := range s {
		// convert current rune to lower case
		r := unicode.ToLower(v)
		// if rune is already in map - string is not unique
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true
}
