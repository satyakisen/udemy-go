package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.golang.org",
		"http://www.stackoverflow.com",
		"http://www.facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLinkStatus(link, c)
		}(l)
	}
}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up and running!")
	c <- link
}
