package channel

import "fmt"

// SendSignalByChannel
// use empty struct or zero length slice of func with channel
// to communicate between go routines
func SendSignalByChannel() {
	ch := make(chan int)
	qCh := make(chan struct{})

	go increment(ch, qCh)

	for i := 0; i < 10; i += 1 {
		fmt.Println(<-ch) // do nothing 10 times
	}
	qCh <- struct{}{} // indicate that the job has done
}

func increment(ch chan int, qCh chan struct{}) {
	var a int
	for {
		select {
		case ch <- a:
			a += 1 // increment by 1 when ch is valid
			fmt.Println(a)
		case <-qCh:
			fmt.Println("end")
			return
		}
	}
}
