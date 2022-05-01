package main

import (
	"fmt"
	"net/http"
	"time"
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

	// for loop for range
	for _, link := range links {
		
		// make http request for each link
		go checkLink(link, c) // only use go keyword infront of function calls
	}

	// for loop - traditional
	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c)
	// }

	// inifinite loop 
	for l := range c{ // this is statement is saying to wait for channel to return a value type and assign it to l, THEN run the body
		// add in a function literal (the same as anonymous function in JS) to sleep the current goroutine
		go func(link string) {
			time.Sleep(10 * time.Second)
			checkLink(link, c)
		}(l) // pass in value l to tell the function literal it will be recieving a parameter, everything in Go is passed by value. The function will create this value in memory and point to it
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // sends a value into channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link // sends a value into channel
}
