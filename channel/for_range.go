package channel

import "fmt"

func ClosableChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i += 1 {
			ch <- i
		}
		close(ch) // if not, deadlock notice will occur
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
