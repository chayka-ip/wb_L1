package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
	собственное множество.
*/

func main() {
	rand.Seed(time.Now().UnixNano())

	/*
		Собственное множетсво любого множества - это такое множество, которое:
		1) не эквивалентно пустому множеству (где нет элементов);
		2) не эквивалентно исходному множеству;
		3) Содержит элемент(ы) из исходного множетсва;
	*/

	v := []string{"cat", "cat", "dog", "cat", "tree"}
	TrivialSolution(v...)
	MapSolution(v...)
}

func TrivialSolution(s ...string) {
	if checkNoFail(s...) {
		i := rand.Intn(len(s))
		fmt.Printf("[%s] is an own set for this set: %v\n", s[i], s)
	}
}

func MapSolution(s ...string) {
	if checkNoFail(s...) {
		// map contains unique keys, so it will be one of the own sets
		res := []string{}
		m := make(map[string]struct{})
		for _, el := range s {
			_, has := m[el]
			if has {
				continue
			}
			m[el] = struct{}{}
			res = append(res, el)
		}
		fmt.Printf("%s is an own set for this set: %v\n", res, s)
	}
}

func checkNoFail(s ...string) bool {
	if len(s) < 2 {
		printFail(s...)
		return false
	}
	return true
}

func printFail(s ...string) {
	fmt.Printf("Own set does not exist for this set: %v\n", s)
}
