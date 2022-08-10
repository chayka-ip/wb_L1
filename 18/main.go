package main

import (
	"fmt"
	"sync"

	c "level_one/18/counter"
)

/*
	Реализовать структуру-счетчик, которая будет
	инкрементироваться в конкурентной среде.
	По завершению программа должна выводить
	итоговое значение счетчика.
*/

var (
	targetValue int64 = 1000
)

func main() {
	// create new counters
	counterAtomic := c.CounterAtomic{}
	counterMutex := c.CounterAtomic{}

	// increment counters in goroutines
	var wg sync.WaitGroup
	for i := 0; i < int(targetValue); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterAtomic.Inc()
		}()
	}

	for i := 0; i < int(targetValue); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterMutex.Inc()
		}()
	}

	wg.Wait()
	LogResult("Atomic", counterAtomic.Get(), targetValue)
	LogResult("Mutex", counterMutex.Get(), targetValue)
}

func LogResult(name string, value int64, target int64) {
	fmt.Printf("%s: %d of %d => ok:%t\n",
		name, value, target, target == value)
}
