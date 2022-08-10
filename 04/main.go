package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

/*
	Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала и
	выводят в stdout. Необходима возможность выбора количества воркеров при
	старте.

	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
	способ завершения работы всех воркеров.
*/

// By default worker count equals to max available parallel goroutines
var defaultWorkersCount = runtime.NumCPU()

func main() {
	var workerCount = getWorkerCount()
	/*
		Context is used to notify goroutines
		that execution must be canceled.

		Сreating new context and cancel callback.
	*/
	ctx, cancel := context.WithCancel(context.Background())
	// wait group for management of goroutines
	wg := &sync.WaitGroup{}

	// Main i/o channel
	mainThread := make(chan int, 100)

	{
		wg.Add(1)
		go runMainThread(ctx, wg, mainThread)
	}

	// launch workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go runWorker(ctx, wg, i, mainThread)
	}

	// new channel to track for CLI input
	interrupt := make(chan os.Signal, 1)
	// listen for interruption or termination commands
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// blocking current function and wait for value from interrupt channel
	<-interrupt
	// unblocking when value is read from interrupt channel
	fmt.Println("\nInterrupting the program...")

	// cancel context to ask goroutines to finish their job
	cancel()

	// wait untill all goroutines are finished
	wg.Wait()
	fmt.Println("Program is finished correctly")
}

func getWorkerCount() int {
	var workerCount int
	// initialize new flag set
	f := flag.NewFlagSet("worker-pool", flag.ContinueOnError)
	// bind variable to argument name
	f.IntVar(&workerCount, "w", defaultWorkersCount, "workers count")
	// parse arguments passed in command line before script was launched
	f.Parse(os.Args[1:])
	if workerCount <= 0 {
		fmt.Printf("Invalid count of workers! Default value (%d) will be used instead\n", defaultWorkersCount)
		// if parsing is invalid - use default value
		workerCount = defaultWorkersCount
	} else {
		fmt.Printf("Number of workers: %d\n", workerCount)
	}
	return workerCount
}

// ch is write-only channle
func runMainThread(ctx context.Context, wg *sync.WaitGroup, ch chan<- int) {
	// remove main tread from WaitGroup
	defer wg.Done()
	for {
		select {
		//if context is canceled => stop main thread
		case <-ctx.Done():
			{
				// close channel and return
				close(ch)
				return
			}
		default:
			// write new random value to the channel
			ch <- int(rand.Intn(1000))
		}
	}
}

// ch is read-only channel
func runWorker(ctx context.Context, wg *sync.WaitGroup, workerId int, ch <-chan int) {
	// remove worker from WaitGroup
	defer wg.Done()
	for {
		select {
		//if context is canceled => stop this worker
		case <-ctx.Done():
			fmt.Printf("Worker %d is exinig because context is closed...\n", workerId)
			return
		/*
			blocked until value is passed through the channel
			or channel is closed
		*/
		case value, ok := <-ch:
			if ok {
				// print new value when have any
				fmt.Printf("Worker %d: %d\n", workerId, value)
			} else {
				//eliminate this worker because channel is closed
				fmt.Printf("Worker %d is exinig because channel is closed...\n", workerId)
				return
			}
		}
	}
}
