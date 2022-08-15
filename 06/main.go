package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	/* 1. Simple return */
	go func() {
		defer wg.Done()
		fmt.Println("✓: Simple return")
	}()

	/*
		2. Tracking for channel closure
	*/
	{
		// data channel
		ch := make(chan struct{})
		wg.Add(2)

		go func() {
			defer wg.Done()
			// send data to the channel
			for i := 0; i < 10; i++ {
				// sending empty struct with zero size
				ch <- struct{}{}
				time.Sleep(1 * time.Millisecond)
			}
			// close channel when data is sent
			close(ch)
		}()

		go func() {
			defer wg.Done()
			/*
				Keep goroutine alive and receive values
				while channel is open
			*/
			for v := range ch {
				_ = v
			}
			fmt.Println("✓: Exit on channel closure")
		}()

	}

	/*
		3. Tracking for context cancelling
	*/
	{
		// Create new context with cancel callback
		ctx, cancel := context.WithCancel(context.Background())
		wg.Add(2)

		go func() {
			defer wg.Done()
			// imitate some work
			time.Sleep(200 * time.Millisecond)
			// cancel context to notify others
			cancel()
		}()

		go func(ctx context.Context) {
			defer wg.Done()
			/*
				Bloked while context is alive.
				Value will be sent from context after cancel
			*/
			<-ctx.Done()
			// Context was closed. Report and exit
			fmt.Println("✓: Exit on context cancelling")
		}(ctx)
	}

	/*
		4. Exit by timer
	*/
	{
		wg.Add(1)
		go func() {
			defer wg.Done()
			t := time.After(300 * time.Millisecond)
			// blocked while some time left
			<-t
			/*
				Time is expired. Data was sent to the timers channel
				and recieved by the statement above
			*/
			fmt.Println("✓: Exit by timer")
		}()
	}

	/*
		5. Exit by signal from trigger channel
	*/
	{
		mainChan, trigger := make(chan struct{}), make(chan struct{})
		wg.Add(3)

		go func() {
			defer wg.Done()
			time.Sleep(400 * time.Millisecond)
			// trigger exit from goroutine
			trigger <- struct{}{}
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				mainChan <- struct{}{}
			}
			close(mainChan)
		}()

		go func() {
			defer wg.Done()
			data := []struct{}{}
			for {
				select {
				case <-trigger:
					{
						fmt.Println("✓: Exit by signal from trigger channel")
						return
					}
				case v, ok := <-mainChan:
					{
						if !ok {
							/*
								Data channel was closed.
								Do some operations while waiting for trigger...
							*/
							_ = len(data)
						} else {
							data = append(data, v)
						}
					}
				}
			}
		}()
	}

	wg.Wait()
}
