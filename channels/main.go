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
		"http://www.freebsd.org",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go getStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			getStatus(link, c)
		}(l)
	}
}

func getStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
