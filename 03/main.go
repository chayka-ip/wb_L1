package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(2^2+3^2+4^2….)
	с использованием конкурентных вычислений.
*/

var a = [...]int32{2, 4, 6, 8, 10}

func main() {
	FirstSolution()
	SecondSolution()
}

func FirstSolution() {

	// sync primitive that traks over goroutine execution
	var wg sync.WaitGroup
	/*
		sync primitive that gives ability
		to prevent multiple acces to some part of code
	*/
	var mu sync.Mutex

	// add number of goroutines to be launched
	wg.Add(len(a))

	// counter
	var sum int32

	for _, x := range a {
		k := x
		go func() {
			defer wg.Done()

			/*
				Mutex is used to prevent multiple access
				to some block of code.
				First goroutine which reaches the Lock instruction
				takes full control over the execution.
				Other goroutines will be blocked at this line.
			*/
			mu.Lock()
			/*
				Goroutine that set up mutex is doing its job
			*/
			sum += k * k
			/*
				Goroutine unlocks mutex
				and next goroutine can take control.
			*/
			mu.Unlock()
		}()
	}

	// wait for each goroutine to finish
	wg.Wait()

	//Output: 220
	fmt.Printf("Result: %d\n", sum)
}

func SecondSolution() {
	var wg sync.WaitGroup

	// add number of goroutines to be launched
	wg.Add(len(a))

	var sum int32

	for _, x := range a {
		k := x
		go func() {
			defer wg.Done()
			/*
				Atomic operation consists of low level instructions
				that don't executed in parallel.
				That means that only one goroutine will have ability
				to change value of passed variable.

				sync package is based on atomic functions
			*/
			atomic.AddInt32(&sum, k*k)
		}()
	}

	// wait for each goroutine to finish
	wg.Wait()

	//Output: 220
	fmt.Printf("Result: %d\n", sum)

}
