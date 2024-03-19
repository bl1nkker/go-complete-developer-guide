package channels

import (
	"fmt"
	"net/http"
	"time"
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
		// Function Literal. It is the same as having "go FunctionLiteralAlternative"
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

		// go some(l, c)
	}
}

func FunctionLiteralAlternative(l string, c chan string){
	time.Sleep(5 * time.Second)
	checkLink(l, c)
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