package http_request

import (
	"fmt"
	"net/http"
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

func SendAsyncReq() {
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

	for _, link := range links {
		link := link
		go func() {
			r, _ := http.Get(link)
			fmt.Println(r)
		}()
	}
}
