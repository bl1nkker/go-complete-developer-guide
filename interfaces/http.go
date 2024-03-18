package interfaces

import (
	"fmt"
	"net/http"
	"os"
)

func RunHttp(){
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}