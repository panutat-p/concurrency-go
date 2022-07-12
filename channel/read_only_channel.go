package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func ListenToChannels() {
	joe := makeSender("🟦")
	ahn := makeSender("🟧")

	// This loop yields 2 channels in sequence
	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ahn)
	}

	// or we can simply use the for range
	// for msg := range joe {
	// 	fmt.Println(msg)
	// }
	fmt.Println("end")
}

// makeSender return read-only channel
func makeSender(msg string) <-chan string {
	ch := make(chan string)
	// we launch goroutine inside a function
	// that sends the data to channel
	go func() {
		// The for loop simulate the infinite sender.
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		// The sender should close the channel
		close(ch)
	}()
	return ch
}
