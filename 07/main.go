package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	mc "level_one/07/map_concurrent"
)

// Реализовать конкурентную запись данных в map.

const (
	numWorkers = 100
)

func main() {
	GenericSolution()
}

func GenericSolution() {
	// create concurrent map instance
	m := mc.NewConcurrentMap()
	var wg sync.WaitGroup

	// launch writing workers to fill the map
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k, v := getRandomInt(), getRandomInt()
			m.Write(k, v)
		}()
	}

	// launch reading workers
	for i := 0; i < numWorkers/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k := getRandomInt()
			m.Get(k)
		}()
	}

	// launch deleting workers
	for i := 0; i < numWorkers/4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k := getRandomInt()
			m.Delete(k)
		}()
	}

	wg.Wait()
}

func SyncMapSolution() {
	/*
		Usecases:
		- (1) when the entry for a given key is only ever written once
			but read many times, as in caches that only grow;
		- (2) when multiple goroutines read, write, and overwrite entries
			for disjoint sets of keys

		It is not recommended to use in 99% cases
	*/
	m := sync.Map{}
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k, v := getRandomInt(), getRandomInt()
			_, loaded := m.LoadOrStore(k, v)
			if loaded {
				fmt.Printf("[%d] already has value:%d\n", k, v)
			} else {
				fmt.Printf("Stored [%d]:%d\n", k, v)
			}
		}()
	}

	wg.Wait()
}

func getRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Intn(100))
}
