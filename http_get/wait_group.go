package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	links = []string{
		"https://amazon.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://example.com",
		"https://golang.org",
		"https://cryptobubbles.net",
		"https://youtube.com",
		"https://googleeee.com",
	}
)

func main() {
	var (
		wg sync.WaitGroup
	)

	for _, link := range links {
		wg.Add(1)
		go get(&wg, link)
	}
	wg.Wait()
}

func get(wg *sync.WaitGroup, link string) {
	defer wg.Done()
	r, err := http.Get(link)
	if err != nil {
		fmt.Println("🟥 err:", err)
		return
	}
	fmt.Println(link, r.Status)
}
