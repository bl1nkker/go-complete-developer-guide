package main

import (
	"fmt"
	"strings"
)

type deck []string

func (d deck) print(){
	fmt.Println("Printing Deck with length", len(d))
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, num int) (deck, deck){
	handDeck := d[:num]
	remainingDeck := d[num:]
	return handDeck, remainingDeck
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

func (d deck) toString() string{
	var result string
	result = strings.Join([]string(d), ",")
	return result
}