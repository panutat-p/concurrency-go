package channel

import "fmt"

func GetFibonacci() {
	ch := make(chan int)
	go writeFibonacci(ch)

	for i := 0; i < 10; i += 1 {
		fmt.Println(<-ch)
	}
}

func writeFibonacci(ch chan int) {
	a, b := 0, 1
	for {
		ch <- a
		a, b = b, a+b
	}
}
