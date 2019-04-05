package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://notrealwebsitejustcrap.com",
	}
	for _, url := range urls {
		go checkUrls(url)
	}

}

func checkUrls(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!")
		return
	}
	fmt.Println(link, "is up!")

}
