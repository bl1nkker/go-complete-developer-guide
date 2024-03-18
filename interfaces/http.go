package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct{}

func (lw LogWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func RunHttp(){
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// bs := make([]byte, 10)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// io.Copy(os.Stdout, resp.Body)
	lw := LogWriter{}
	io.Copy(lw, resp.Body)
}