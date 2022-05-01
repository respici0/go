package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // how to create a channel of a given type, in this case string

	for _, link := range links {
		// make http request for each link
		go checkLink(link, c) // only use go keyword infront of function calls
	}

	fmt.Println(<-c)
	fmt.Println(<-c) // function call to send value of channel
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think" // send a value into channel
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep it's up" // send a value into channel
}
