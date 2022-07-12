package channel

import (
	"fmt"
	"net/http"
)

func SendAsyncReqChannel() {
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

	ch := make(chan string)

	for _, link := range links {
		link := link
		go func() {
			r, err := http.Get(link)
			if err != nil {
				ch <- err.Error()
				return
			}
			ch <- fmt.Sprintf("%v %v", link, r.Status)
		}()
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch) // block
	}
}
