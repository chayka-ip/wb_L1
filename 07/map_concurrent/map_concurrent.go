package map_concurrent

import (
	"sync"
)

type ConcurrentMap struct {
	/*
		Not any type is allowed to be key of map.
		It is up to developers to use map with proper data.
	*/
	data  map[any]any
	mutex *sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data:  make(map[any]any),
		mutex: &sync.RWMutex{},
	}
}

// Adds or overwrites value for the key given
func (m *ConcurrentMap) Write(key any, value any) {
	/*
		Set read-write lock to prevent multiple access while writing.
		Lock can be set only when there are no otner Locks or RLocks.
		In other words, only one read-write lock is possible.
		If there are some RLocks present - blocked until they released,
		after that - set Lock
	*/
	m.mutex.Lock()
	// release lock after data was updated
	defer m.mutex.Unlock()
	// update data
	m.data[key] = value
}

/*
	If value is found for key => returns value and true
	Otherwise returns default value and false
*/
func (m *ConcurrentMap) Get(key any) (any, bool) {
	/*
		Read lock can be set only when there is no Lock (read-write lock).
		There can be multiple Read locks at the time.
		If Lock is set - blocked until it is released
	*/
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

// Deletes given key if present
func (m *ConcurrentMap) Delete(key any) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, key)
}

// True if key in the map, false if not
func (m *ConcurrentMap) Has(key any) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	_, ok := m.data[key]
	return ok
}
