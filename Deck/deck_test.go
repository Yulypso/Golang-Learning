package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSave2FileAndTestNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.save2File("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 52 {
		t.Errorf("Expected loaded deck length of 52 but got %v", len(loadDeck))
	}

	os.Remove("_decktesting")
}
