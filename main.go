package main

import (
	"go-complete-developer-guide/cards"

)

func main() {
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