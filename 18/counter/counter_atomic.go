package counter

import "sync/atomic"

/*
	Counter with atomic-based functions.

	Atiomic provides low-level operations
	that exclude multiple access to the memory.

	Atomic works times faster than mutex
	when dealing with such simple operations
	like read/write numbers etc...

*/
type CounterAtomic struct {
	value int64
}

func (c *CounterAtomic) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *CounterAtomic) Get() int64 {
	return atomic.LoadInt64(&c.value)
}
