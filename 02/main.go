package main

import (
	"fmt"
	"sync"
)

var (
	a           = [...]int{2, 4, 6, 8, 10}
	workerCount = 3
)

/*
	Написать программу, которая конкурентно рассчитает
	значение квадратов чисел взятых из массива (2,4,6,8,10)
	и выведет их квадраты в stdout.
*/

func main() {
	firstSolution()
	fmt.Println("==============")
	secondSolution()
}

func firstSolution() {
	// declare wait group to proceed all required tasks
	var wg sync.WaitGroup
	// set number of tasks
	wg.Add(len(a))

	// iterate over data array
	for _, x := range a {
		// run anonymous function with buisness logic as goroutine
		go func(k int) {
			// close task when it's done
			defer wg.Done()
			//buisness logic
			fmt.Printf("%d * %d = %d\n", k, k, k*k)
		}(x) //x is passed to function as argument to work with
	}

	// wait untill all tasks are done
	wg.Wait()
}

func secondSolution() {
	// make channel to write data to
	ch := make(chan int, len(a))
	var wg sync.WaitGroup

	// run writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range a {
			ch <- v
		}
		close(ch)
	}()

	// run workers to process data from the channel
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go runWorker(&wg, ch)
	}

	wg.Wait()
}

// ch is read-only channel
func runWorker(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	// while there is some data in the channel ant it is open
	for k := range ch {
		// process data
		fmt.Printf("%d * %d = %d\n", k, k, k*k)
	}
}
