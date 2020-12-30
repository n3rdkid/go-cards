package main

func main() {
	cards := newDeckFromFile("eck_01")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
