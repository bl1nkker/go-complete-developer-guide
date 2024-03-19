package channels

import (
	"fmt"
	"net/http"
	"time"
)

func RunChannels(){
	fmt.Println("--------------Channels--------------\n\n\n\n\n")
	links := []string{
		"http://google.com", 
		"http://golang.org", 
		"http://facebook.com",
		"http://amazon.com",
	}

	// Create channel to communicate between goroutines
	c := make(chan string)

	for i, link := range links{
		// HTTP Request. Create Goroutine (thread)
		fmt.Println("F1:: Iteration number - ", i, "with link:", link)
		message := "--For Loop 1::Goroutine " + fmt.Sprint(i) + ":: data " + link
		go checkLink(link, c, message)
	}
	// Blocking Line of code. The program will listen the channel until it receives some messages
	fmt.Println(<- c)
	counter := 0
	for l := range c{

		fmt.Println("F2:: Iteration number - ", counter, "with link:", l)
		// Function Literal. It is the same as having "go FunctionLiteralAlternative"
		go func(link string, counter int) {
			time.Sleep(5 * time.Second)
			message := "--For Loop 2::Goroutine " + fmt.Sprint(counter) + ":: data " + l
			checkLink(link, c, message)
		}(l, counter)
		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	message := "--For Loop 2::Goroutine " + cfmt.Sprint(counter) + ":: data " + l
		// 	checkLink(l, c, message)
		// }()
		counter ++

		// go some(l, c)
	}
}

func FunctionLiteralAlternative(l string, c chan string){
	time.Sleep(5 * time.Second)
	checkLink(l, c, "")
}

func checkLink(link string, c chan string, message string){
	// Blocking Line of code
	fmt.Println(message, "START")
	http.Get(link)
	c <- link
	fmt.Println(message, "END")
}