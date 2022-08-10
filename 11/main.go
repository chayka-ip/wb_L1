package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

//Реализовать пересечение двух неупорядоченных множеств.

const (
	MIN_SLICE_SIZE = 6
	MAX_SLICE_SIZE = 10
	MAX_VALUE      = 9
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := getRandomIntSlice(), getRandomIntSlice()
	ru := IntersectUnique(a, b)
	rf := IntersectFull(a, b)
	fmt.Printf("%v | %v => unique %v\n", a, b, ru)
	fmt.Printf("%v | %v => full %v\n", a, b, rf)
}

/*
	Returns intersection where all elements are unique

	Ordered is a generic interface.
	Than mean that function can be used for diferent types
	supporting comparison operations.
*/
func IntersectUnique[T constraints.Ordered](a []T, b []T) []T {
	out := []T{}
	m := make(map[T]int)
	for _, k := range a {
		_, has := m[k]
		if has {
			continue
		}
		m[k] = 1
	}
	for _, k := range b {
		v, has := m[k]
		if has {
			if v == 1 {
				m[k] += 1
				out = append(out, k)
			}
		}
	}
	return out
}

/*
	Returns full intersection of two sets (i.e. the most larger set
	that can be found in both a and b sets.
*/
func IntersectFull[T constraints.Ordered](a []T, b []T) map[T]int {
	out := make(map[T]int)
	m1, m2 := sliceToMap(a), sliceToMap(b)

	for k, v1 := range m1 {
		if v2, has := m2[k]; has {
			out[k] = Min(v1, v2)
		}
	}

	return out
}

// returns map of keys and count of that key in the map
func sliceToMap[T constraints.Ordered](s []T) map[T]int {
	m := make(map[T]int)
	for _, v := range s {
		m[v] += 1
	}
	return m
}

func Min[T constraints.Ordered](a T, b T) T {
	if a >= b {
		return b
	}
	return a
}
func getRandomIntSlice() []int {
	size := rand.Intn(MAX_SLICE_SIZE-MIN_SLICE_SIZE+1) + MIN_SLICE_SIZE
	out := make([]int, size)
	for i := range out {
		out[i] = rand.Intn(MAX_VALUE + 1)
	}
	return out
}
