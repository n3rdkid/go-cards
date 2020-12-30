package main

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 4)
	hand.print()
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
