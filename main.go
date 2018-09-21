package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://facebook.com",
		"https://google.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// since go is a pass by value language
		// we're going to send in the link
		// value as a property, so the child routine
		// has a copied value of the link
		go func(link string) { // `function literal`
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
