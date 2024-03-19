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
	for _, link := range links{
		// HTTP Request
		checkLink(link)
	}
}

func checkLink(link string) bool{
	_, err := http.Get(link)
	if err != nil{
		fmt.Println("Link might be down:", link)
		return false
	}
	fmt.Println("Link is up:", link)
	return true
}