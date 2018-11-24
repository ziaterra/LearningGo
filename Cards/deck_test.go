package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 4*4 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card ot be Ace of Spades")
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of clubs")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	// delete any files
	os.Remove("_decktesting")

	// create a deck
	d := newDeck()
	fn := "_decktesting"

	// Save file called "_decktesting"
	d.saveToFile(fn)

	// Load from File
	newd := newDeckFromFile(fn)

	// Assert deck length
	if len(newd) != 16 {
		t.Errorf("Expected len of 16 but got %v", len(newd))
	}

	// delete any files in current working directory with the name "_decktesting"
	os.Remove(fn)
}
