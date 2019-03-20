package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
        // list of urls that we want to monitor
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)
        
        // shoot a gorutine for each url
	for _, url := range urls {
		go checkURL(url, c)
	}
        
	for l := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkURL(url, c)
		}(l)
	}
}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
