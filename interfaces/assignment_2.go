package interfaces

import (
	"fmt"
	"io"
	"os"
)

func RunAssignment2(){
	// Get filename
	if len(os.Args) < 1{
		fmt.Println("Filename must be specified... Cinamononon rouuul")
		os.Exit(1)
	}
	fw, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("Error occured when reading file:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, fw)
}