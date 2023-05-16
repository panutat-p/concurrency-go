package main

import "fmt"

func main() {
	var (
		c = make(chan int)
	)
	go calculate(c)

	for i := 0; i < 10; i += 1 {
		fmt.Println(<-c) // read
	}
}

func calculate(ch chan int) {
	var (
		a = 0
		b = 1
	)
	for {
		ch <- a // write
		a, b = b, a+b
	}
}
