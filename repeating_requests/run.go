package repeating_requests

import (
	"fmt"
	"net/http"
	"time"
)

// pass the channel into this function
// and store proceeded value into the channel
func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link // send link back to main routine
		return
	}
	fmt.Println(link, "is up!")
	ch <- link // send link back to main routine
}

func Run() {
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

	for l := range ch { // each link in slice above
		go func(lr string) {
			time.Sleep(5 * time.Second) // execute all at once then wait 5s
			checkLink(lr, ch)           // avoid using "l" because of closure
		}(l) //  // pass each link to link for each routine
	}
}
