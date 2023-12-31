package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "A of Spades" {
		t.Errorf("Expected first card of A of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expected first card of K of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadeDeck := newDeckFromFile(("_decktesting"))

	if len(loadeDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadeDeck))
	}

	os.Remove("_decktesting")
}
