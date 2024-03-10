package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	// Returns a new slice
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string{
	return "Five of Diamonds"
}