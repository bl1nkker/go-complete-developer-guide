package main

import (
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
	alexContact := structs.ContactInfo{Email: "alex@gmail.com", Number: "111333"}
	alex.FirstName = "Alex"
	alex.LastName = "Anderson"
	alex.ContactInfo = alexContact
	alex.Print()

	jim := structs.Person{
		FirstName: "Jim",
		LastName: "Fatherbeard",
		ContactInfo: structs.ContactInfo{
			Email: "jim@gmail.com",
			Number: "228337",
		},
	}
	jim.UpdateName("Fuck")
	jim.Print()
}