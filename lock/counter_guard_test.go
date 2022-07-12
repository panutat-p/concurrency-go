package lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter_Increment(t *testing.T) {
	c := &Counter{
		mu:    &sync.Mutex{},
		value: 0,
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
			fmt.Print(".")
		}()
	}
	wg.Wait()
	fmt.Println("\ndone")

	if c.value != 100 {
		t.Error("want 100 but got", c.value)
	}
}
