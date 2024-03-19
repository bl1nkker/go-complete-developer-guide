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
	fmt.Println(<- c)
}

func checkLink(link string, c chan string){
	_, err := http.Get(link)
	if err != nil{
		fmt.Println("Link might be down:", link)
		c <- "Might be down, i think"
		return
	}
	fmt.Println("Link is up:", link)
	c <- "Yep, it is up"	
}