package main

import (
	"fmt"
	"net/http"
	"time"
)

// status checker
func main() {
	urls := []string{
		"http://google.com.sg",
		"http://amazon.com.sg",
		"http://facebook.com.sg",
		"http://shopify.com.sg",
	}

	c := make(chan string)

	// 	initialize channels
	for _, url := range urls {
		go pingSite(url, c)
	}

	// Watch the channel c
	// If there is a new value, we fetch a link.
	// i.e. "recieve message from the channel"
	for l := range c { // using range i-th a channel. Wait for a channel
		go func(link string) { // scope of l must be passed into function literal. Pass by value.
			time.Sleep(time.Second * 5) // pause for 5 seconds
			pingSite(link, c)
		}(l) // l is copied in memory, go routine has access to copy.
	}
}

func pingSite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Url: " + url + " is down.")
		c <- "Might be down"
		fmt.Println(err)
		return
	}
	// Log
	fmt.Println("Url: " + url + " is up!")

	// Send message in the channel
	c <- url
}
