package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

/**
Creates a new deck
*/
func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/**
Deal cards from the deck
@param deck : The deck from which we deal
@param handSize: The no of cards to deal
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/**
Prints the current cards in the deck
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/**
Turn a deck into a string
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/**
Save a deck into a file
@param filename: Name of the file where the deck will be stored
*/
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

/**
Convert a string into deck
@param filename: Name of the file from the deck will be loaded
*/
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

/**
Shuffles the deck
*/
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
