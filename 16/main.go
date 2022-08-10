package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

/*
	Реализовать быструю сортировку массива (quicksort)
	встроенными методами языка.
*/

const (
	MIN_SLICE_SIZE = 10
	MAX_SLICE_SIZE = 10
	MIN_VALUE      = -10
	MAX_VALUE      = 10
)

func main() {
	d := getRandomIntSlice()
	fmt.Printf("Before: %d\n", d)
	QuickSort(d)
	fmt.Printf("After: %d\n", d)
}

// sorts slice in ascending order
func QuickSort[T constraints.Ordered](slice []T) {
	// There is nothing to sort in empty slice or when size is 1
	if len(slice) <= 1 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go quickSortBody(&wg, slice, 0, len(slice)-1)
	wg.Wait()
}

// sorts slice in ascending order
func quickSortBody[T constraints.Ordered](wg *sync.WaitGroup, slice []T, startInd int, endInd int) {
	defer wg.Done()
	if startInd >= endInd {
		return
	}

	// define pivot index and find correspondng element
	pivotInd := (endInd + startInd) / 2
	pivotElem := slice[pivotInd]

	// define indieces to iterate through elements from ends
	i, j := startInd, endInd

	for i <= j {

		for {
			if slice[i] < pivotElem {
				i++
			} else {
				/*
					Exit when there is element on the left side
					which is greater then or equal to pivot value
					or pivot is reached
				*/
				break
			}
		}

		for {
			if slice[j] > pivotElem {
				j--
			} else {
				/*
					Exit when there is element on the right side
					which is less then or equal to pivot value
					or pivot is reached
				*/
				break
			}
		}

		if i <= j {
			// swap elements
			slice[i], slice[j] = slice[j], slice[i]
			// shift indices to define new bounds
			i++
			j--
		}
	}

	// process further ranges in parallel
	wg.Add(2)
	go quickSortBody(wg, slice, startInd, j)
	go quickSortBody(wg, slice, i, endInd)
}

func getRandomIntSlice() []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(MAX_SLICE_SIZE-MIN_SLICE_SIZE+1) + MIN_SLICE_SIZE
	out := make([]int, size)
	for i := range out {
		amv := int(math.Abs(MIN_VALUE))
		mv := MAX_VALUE + amv + 1
		out[i] = rand.Intn(mv) - amv
	}
	return out
}
