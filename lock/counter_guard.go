package lock

import "sync"

/*
When we want to guard internal state of a struct
use memory access synchronization

👍 Hide detail of locking from the caller
😃 Try to keep the locks constrained to a small lexical scope
*/

type Counter struct {
	mu    *sync.Mutex
	value int
}

// Increment
// call to Increment can be considered atomic
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}
