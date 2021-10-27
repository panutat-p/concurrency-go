package main

import (
	"fmt"
	"net/http"
)

// pass the channel into this function
// and store proceeded value into the channel
func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- "Might be down" // store value to channel
		return
	}
	fmt.Println(link, "is up!")
	ch <- "It's up" // store value to channel
}

func main() {
	links := []string{
		"https://amazon.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://example.com",
		"https://golang.org",
		"https://cryptobubbles.net",
		"https://youtube.com",
		"https://google.com",
	}

	ch := make(chan string) // create go channel to store string

	for _, link := range links {
		go checkLink(link, ch)
	}

	// blocking channel, if main routine got a value, it will exit
	fmt.Println(<-ch) // wait for a value, when get one, print it out
}
