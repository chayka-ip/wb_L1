package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"golang.org/x/exp/constraints"
)

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

const (
	MIN_SLICE_SIZE = 10
	MAX_SLICE_SIZE = 20
	MIN_VALUE      = -10
	MAX_VALUE      = 10
)

func main() {
	rand.Seed(time.Now().UnixNano())
	/*
		Value to search.
		Might be not presented in the actual data.
	*/
	searchVal := getRandomValue()
	data := getRandomIntSlice()
	// data must be sorted before binary search
	sort.Ints(data)

	index, found := BinarySearch(data, searchVal)

	s := "not found"
	if found {
		s = fmt.Sprintf("index: %d", index)
	}
	fmt.Printf("Source data: %v\n", data)
	fmt.Printf("Search value: %d | %s\n", searchVal, s)

}

/*
	Constrains allows this function work with primitive types
	that support such operations as < <= >= > (Integer | Float | ~string)
*/
func BinarySearch[T constraints.Ordered](slice []T, val T) (int, bool) {
	// define not found result to reuse
	nf := func() (int, bool) { return -1, false }
	// define bound indices
	iLow, iHigh := 0, len(slice)-1

	for iLow <= iHigh {
		// if value not in the bounds - it is not presented
		if (val < slice[iLow]) || (val > slice[iHigh]) {
			return nf()
		}
		// get middle index
		mid := (iLow + iHigh) / 2
		if val == slice[mid] {
			return mid, true
		}
		// update bounds
		if val > slice[mid] {
			iLow = mid + 1
		} else {
			iHigh = mid - 1
		}
	}

	return nf()
}

func getRandomValue() int {
	amv := int(math.Abs(MIN_VALUE))
	mv := MAX_VALUE + amv + 1
	return rand.Intn(mv) - amv
}

func getRandomIntSlice() []int {
	size := rand.Intn(MAX_SLICE_SIZE-MIN_SLICE_SIZE+1) + MIN_SLICE_SIZE
	out := make([]int, size)
	for i := range out {
		out[i] = getRandomValue()
	}
	return out
}
