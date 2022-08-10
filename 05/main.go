package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в
	канал, а с другой стороны канала — читать. По истечению N секунд программа
	должна завершаться.
*/

func main() {
	// create channel
	ch := make(chan struct{})
	// set work time
	t := 1 * time.Second

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		/*
			Timer is initialized once outside of the loop
		*/
		et := time.After(t)
		for {
			select {
			// if time is out => value is extracted from timer channel
			case <-et:
				// close channel and exit from goroutine
				close(ch)
				fmt.Println("Time is over. Closing channel...")
				return
			// we write data to the channel while there is time left
			default:
				ch <- struct{}{}
			}
		}
	}()

	go func() {
		defer wg.Done()
		// iterate while channel remains open
		for i := 1; ; i++ {
			_, ok := <-ch
			// exit if channel is closed
			if !ok {
				fmt.Println("Exting from reader...")
				return
			}
			// print data when new value obtained
			fmt.Printf("%d: ...\n", i)
		}
	}()

	// wait while goroutines do their job
	wg.Wait()
	fmt.Println("OK")
}
