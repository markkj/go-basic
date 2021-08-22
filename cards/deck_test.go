package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card of Five of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := loadFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck, got %v", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", loadedDeck[0])
	}

	if loadedDeck[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card of Five of Spades, but got %v", loadedDeck[len(d)-1])
	}

	os.Remove("_decktesting")
}
