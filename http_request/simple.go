package http_request

import (
	"fmt"
	"net/http"
	"sync"
)

func SendSyncReq() {
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

	for _, v := range links {
		r, _ := http.Get(v)
		fmt.Println(v, r.Status)
	}
}

func SendAsyncReqChannel() {
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

	ch := make(chan string)

	for _, link := range links {
		link := link
		go func(ch chan string) {
			r, _ := http.Get(link)
			ch <- fmt.Sprintf("%v %v", link, r.Status)
		}(ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch) // block
	}
}

func SendAsyncReqRoutine() {
	wg := &sync.WaitGroup{}

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

	wg.Add(len(links))

	for _, link := range links {
		link := link
		go func(wg *sync.WaitGroup) {
			r, _ := http.Get(link)
			fmt.Println(link, r.Status)
			wg.Done()
		}(wg)
	}
	wg.Wait() // block
}
