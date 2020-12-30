package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck length of 16, but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubes, but got %v", d[len(d)-1])
	}

}

func TestSaveToDectAndNewDeckFromFile(t *testing.T) {
	// Perform Cleanup
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	deckFromFile := newDeckFromFile("_decktesting")

	if len(deckFromFile) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(deckFromFile))
	}
}
