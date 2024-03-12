package main

func main() {
	cards := newDeck()
	hand, cards := deal(cards, 5)
	hand.print()
	cards.print()
	cards.saveToFile("cards.txt")
	newD := newDeckFromFile("cards.txt")
	newD.print()
}