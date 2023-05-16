package main

import "fmt"

func main() {
	var (
		ch = make(chan int)
	)

	go do(ch)

	for e := range ch {
		fmt.Println(e)
	}
	// exit the loop when the channel is closed
}

func do(ch chan<- int) {
	for i := 1; i < 10; i += 1 {
		ch <- i
	}
	close(ch) // 🦊 prevent deadlock
}
