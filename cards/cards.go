package cards

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Deck []string

func (d Deck) Print(){
	fmt.Println("Printing Deck with length", len(d))
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d Deck, num int) (Deck, Deck){
	handDeck := d[:num]
	remainingDeck := d[num:]
	return handDeck, remainingDeck
}

func NewDeck() Deck{
	cards := Deck{}
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

func (d Deck) ToString() string{
	result := strings.Join([]string(d), ",")
	return result
}

func (d Deck) SaveToFile(filename string) error {
	data := []byte(d.ToString())
	res := os.WriteFile(filename, data, 0666)
	return res
}

func NewDeckFromFile(filename string) Deck{
	bs, err := os.ReadFile(filename)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return Deck(strings.Split(string(bs), ","))
}

func (d Deck) Shuffle() {
	for i := range d{
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func Run(){
	deck := NewDeck()
	_, deck = Deal(deck, 5)

	// I/O
	deck.SaveToFile("cards.txt")
	newD := NewDeckFromFile("cards.txt")
	newD.Print()

	// Shuffle
	deck.Print()
	deck.Shuffle()
	deck.Print()
}