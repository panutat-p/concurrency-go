package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		MaxJob    = 9
		MaxWorker = 3
	)

	var (
		jobs    = make(chan int, MaxJob)
		results = make(chan int, MaxJob)
	)

	// Start workers
	for i := 1; i <= MaxWorker; i++ {
		go start(i, jobs, results)
	}

	// Send jobs to workers
	for i := 1; i <= MaxJob; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	for i := 1; i <= MaxJob; i++ {
		r := <-results
		fmt.Println("🍌", r)
	}
}

func start(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println("🐵 worker", id, "job", i)
		time.Sleep(time.Second)
		results <- i * 100
	}
}
