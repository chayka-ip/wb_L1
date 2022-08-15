package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {

	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(removeItemBreakOrder(slice, 3)) // Output: [0 1 2 5 4]
	fmt.Println(slice)                          // Output: [0 1 2 3 4 5]

	fmt.Println(removeItemBreakOrderInplace(slice, 3)) // Output: [0 1 2 5 4]
	fmt.Println(slice)                                 // Output: [0 1 2 5 4 5]

	slice = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(removeItemPreserveOrder(slice, 3)) // Output: [0 1 2 4 5]
	fmt.Println(slice)                             // Output: [0 1 2 4 5 5]

	slice = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(removeItemByCopy(slice, 3)) // Output: [0 1 2 4 5]
	fmt.Println(slice)                      // Output: [0 1 2 4 5 5]
}

/*
	Removes i-th slice element.
	Modifies source slice.
	Valid input is considered.
	[T any] means that function will work with slice of any type
*/
func removeItemPreserveOrder[T any](slice []T, i uint64) []T {
	return append(slice[:i], slice[i+1:]...)
}

/*
	Removes i-th slice element.
	Modifies source slice.
	Valid input is considered.
*/
func removeItemBreakOrderInplace[T any](slice []T, i uint64) []T {
	lastInd := len(slice) - 1
	slice[i] = slice[lastInd]
	return slice[:lastInd]
}

/*
	Removes i-th slice element.
	Does not modify source slice.
	Valid input is considered.
*/
func removeItemBreakOrder[T any](slice []T, i uint64) []T {
	lastInd := len(slice) - 1
	s := append([]T{}, slice...)
	s[i] = slice[lastInd]
	return s[:lastInd]
}

/*
	Removes i-th slice element inplace.
	Modifies source slice.
	Valid input is considered.
*/
func removeItemByCopy[T any](slice []T, i uint64) []T {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
