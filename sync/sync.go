package sync

import "sync"

// Counter is used to hold the counting value
type Counter struct {
	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	mu    sync.Mutex
	value int
}

// Inc is implemeted to increase the counter to 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value is implemented to get the current counter value
func (c *Counter) Value() int {
	return c.value
}
