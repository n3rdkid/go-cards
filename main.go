package main

func main() {
	cards := newDeck()
	cards.saveToFile("deck_01")
}

func newCard() string {
	return "Five of Diamonds"
}
