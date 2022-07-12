package channel

import "fmt"

/*
fork 1 go routine
*/

func GetFibonacci() {
	ch := make(chan int)
	go infiniteFibo(ch)

	for i := 0; i < 10; i += 1 {
		fmt.Println(<-ch) // read fibo sequence from chanel
	}
}

func infiniteFibo(ch chan int) {
	a, b := 0, 1 // go routine hold these value
	for {
		ch <- a // write fibo sequence to chanel
		a, b = b, a+b
	}
}
