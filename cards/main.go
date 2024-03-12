package main

func main() {
	cards := newDeck()
	_, cards = deal(cards, 5)

	// I/O
	cards.saveToFile("cards.txt")
	newD := newDeckFromFile("cards.txt")
	newD.print()

	// Shuffle
	cards.print()
	cards.shuffle()
	cards.print()
}