package counter

import (
	"sync"
)

/*
	Counter with mutex-based functions.

	Mutex grants unique access to some code instuctions.
	It is usefull when few operations should be
	executed together but data race might occur.
*/

type CounterMutex struct {
	value int64
	mutex *sync.RWMutex
}

func (c *CounterMutex) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *CounterMutex) Get() int64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}
