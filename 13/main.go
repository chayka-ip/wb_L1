package main

import (
	"fmt"
)

/*
	Поменять местами два числа без создания временной переменной.
*/

func main() {
	a, b := 10, 100

	// Output: a = 10 | b = 100
	fmt.Printf("a = %d | b = %d\n", a, b)

	a, b = b, a

	// Output: a = 100 | b = 10
	fmt.Printf("a = %d | b = %d\n", a, b)
}
