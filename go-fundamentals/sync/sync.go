package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// any goroutine calling Inc will acquire the lock on Counter if they are first.
	// All the other goroutines will have to wait for it to be Unlocked before getting access.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

/*
go vet
Remember to use go vet in your build scripts as it can alert you to some subtle bugs
in your code before they hit your poor users.
*/
