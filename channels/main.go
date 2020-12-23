package main

import (
	"fmt"
	"net/http"
	"time"
)

var links = [...]string{"https://www.google.com",
	"https://www.facebook.com",
	"https://www.whatsapp.com",
	"https://www.hepsiburada.com",
	"https://www.trendyol.com",
	"https://stackoverflow.com"}

func main() {

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// for l := range c {
	// 	go func(l string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(l, c)
	// 	}(l)
	// }

}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down")
		time.Sleep(5 * time.Second)
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	time.Sleep(5 * time.Second)
	c <- link

}
