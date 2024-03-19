package channels

import (
	"fmt"
	"net/http"
)

func RunChannels(){
	links := []string{
		"http://google.com", 
		"http://golang.org", 
		"http://facebook.com",
		"http://amazon.com",
	}

	// Create channel to communicate between goroutines
	c := make(chan string)

	for _, link := range links{
		// HTTP Request. Create Goroutine (thread)
		go checkLink(link, c)
	}
	// Blocking Line of code. The program will listen the channel until it receives some messages
	fmt.Println(<- c)

	for l := range c{
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string){
	// Blocking Line of code
	_, err := http.Get(link)
	if err != nil{
		fmt.Println("Link might be down:", link)
		c <- link
		return
	}
	fmt.Println("Link is up:", link)
	c <- link
}