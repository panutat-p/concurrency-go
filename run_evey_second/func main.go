package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		c1 = make(chan string)
		c2 = make(chan string)
	)

	go func() {
		for {
			c1 <- "every 1s"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case <-c1:
			fmt.Println("🐢")
		case <-c2:
			fmt.Println("🐵")
		}
	}
}
