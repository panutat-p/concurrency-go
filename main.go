package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var c <-chan interface{}
	wg := sync.WaitGroup{}
	noop := func() {
		wg.Done()
		<-c // go routine will never exit, so we can keep them in memory
	}

	const numGoRoutines = 1e4
	wg.Add(numGoRoutines)

	before := MemConsumed()

	for i := numGoRoutines; i > 0; i -= 1 {
		go noop()
	}
	wg.Wait()

	after := MemConsumed()

	fmt.Printf("%.3f KB", float64(after-before)/numGoRoutines/1000)
}

func MemConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
