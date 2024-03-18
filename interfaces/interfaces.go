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
type FrenchBot struct{}

func (eb EnglishBot) GetGreeting() string{
	return "Hello!"
}

func (jb JapaneseBot) GetGreeting() string{
	return "Arigato!"
}

func (fb FrenchBot) GetGreeting() string {
    return "Bonjour!"
}

func Run(){
	eb := EnglishBot{}
	jb := JapaneseBot{}
	fb := FrenchBot{}
	PrintGreeting(eb)
	PrintGreeting(jb)
	PrintGreeting(fb)
}