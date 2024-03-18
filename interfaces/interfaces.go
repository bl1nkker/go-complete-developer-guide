package interfaces

import "fmt"

type bot interface{
	GetGreeting() string
}

func PrintGreeting(b bot){
	fmt.Println(b.GetGreeting())
}

type EnglishBot struct{}
type JapaneseBot struct{}

func (eb EnglishBot) GetGreeting() string{
	return "Hello!"
}

func (jb JapaneseBot) GetGreeting() string{
	return "Arigato!"
}