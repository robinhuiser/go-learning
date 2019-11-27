package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be an Ace of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected first card to be a Four of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	tmpFileName := "_decktesting"

	// Remove any left over files from previous tests
	// Should be moved to Makefile in future I guess...
	os.Remove(tmpFileName)

	// Instantiate new deck
	d := newDeck()
	d.saveToFile(tmpFileName)

	// Load new deck from same file
	df := newDeckFromFile(tmpFileName)

	// Assertions
	if len(df) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(df))
	}

	// Cleanup
	os.Remove(tmpFileName)
}
