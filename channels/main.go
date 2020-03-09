package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.golang.org",
		"http://www.stackoverflow.com",
		"http://www.facebook.com",
	}

	for _, link := range links {
		checkLinkStatus(link)
	}
}

func checkLinkStatus(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!")
		return
	}

	fmt.Println(link, "is up and running!")
}
