package sleep

import (
	"fmt"
	"time"
)

func PrintForOneSecond() {
	go printMinus()
	go printPlus()
	time.Sleep(1 * time.Second) // block for 1 second
}

func printMinus() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("-")
	}
}

func printPlus() {
	for {
		time.Sleep(200 * time.Millisecond)
		fmt.Print("+")
	}
}
