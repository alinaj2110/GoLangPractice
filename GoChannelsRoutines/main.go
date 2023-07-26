package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://youtube.com",
		"http://wronglink",
	}

	c := make(chan string)
	for _, link := range links {
		go checklink(link, c)
	}

	for l := range c {
		time.Sleep(5 * time.Second)
		go func(linkcopy string) {
			checklink(linkcopy, c)
		}(l)
	}
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down.")
		c <- link
		return
	}

	fmt.Println(link, "is up.")
	c <- link
}
