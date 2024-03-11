package main

import "fmt"

type deck []string

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck{
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Ace"}
	var newCard string
	for _, suit := range cardSuits{
		for _, val := range cardValues{
			newCard = val + " of " + suit
			cards = append(cards, newCard)
		}
	}
	return cards
}