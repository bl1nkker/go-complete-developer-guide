package main

import (
	"fmt"
	"go-complete-developer-guide/cards"
	"go-complete-developer-guide/structs"
)

// Cards
func main() {
	// runCards()
	runStructs()
}

func runCards(){
	deck := cards.NewDeck()
	_, deck = cards.Deal(deck, 5)

	// I/O
	deck.SaveToFile("cards.txt")
	newD := cards.NewDeckFromFile("cards.txt")
	newD.Print()

	// Shuffle
	deck.Print()
	deck.Shuffle()
	deck.Print()
}

func runStructs(){
	// alex := structs.Person{FirstName: "Alex", LastName: "Anderson"}
	var alex structs.Person
	alex.FirstName = "Alex"
	alex.LastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}