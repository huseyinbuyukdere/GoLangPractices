package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got  %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but god %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but god %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Loaded deck expected 16 but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
