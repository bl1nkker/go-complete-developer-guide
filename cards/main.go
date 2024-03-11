package main

func main() {
	cards := newDeck()
	hand, cards := deal(cards, 5)
	hand.print()
	cards.print()

}