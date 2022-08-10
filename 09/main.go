package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
	массива, во второй — результат операции x*2, после чего данные из второго
	канала должны выводиться в stdout.
*/

var (
	// source data array
	data = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// number of workers
	nw = getNumWorkers()
)

func main() {
	// channels
	cin, cout := make(chan int, len(data)), make(chan int)

	/*
		As far as write operation to cout occurs in multiple goroutines
		we should close this channel from there.

		Close function may be called only once (it panics if called more times).

		sync.Once mutex is used to prevent multiple attempts to close the channel.
		Second and consequent calls would be no-op if mutex is already activated.
	*/
	var closeOutChanOnce sync.Once
	closeOutChan := func() { close(cout) }

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// write from data array to input channel
		for _, v := range data {
			cin <- v
		}
		// close input channel
		close(cin)
	}()

	// create a few workers to square input numbers
	for i := 0; i < nw; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// write data to out channel while input is present
			for v := range cin {
				cout <- v * v
			}
			/*
				Close output channel.
				closeOutChan is executed only once.
			*/
			closeOutChanOnce.Do(closeOutChan)
		}()
	}

	// run some workers to print output
	for i := 0; i < nw; i++ {
		wg.Add(1)
		go runPrintWorker(wg, cout)
	}

	wg.Wait()
}

func runPrintWorker(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("Recieved: %d\n", v)
	}
}

func getNumWorkers() int {
	n := runtime.NumCPU()
	if n > 1 {
		n /= 2
	}
	return n
}
