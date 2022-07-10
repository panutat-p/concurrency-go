package http_request

import (
	"fmt"
	"log"
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
		"https://googles.com",
	}

	wg.Add(len(links))

	for _, link := range links {
		link := link
		go func() {
			defer wg.Done()
			r, err := http.Get(link)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(link, r.Status)
		}()
	}
	wg.Wait() // block
}
