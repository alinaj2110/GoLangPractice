package main

import (
	"os"
	"testing"
)

func TestDeckLength(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52, got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveFileandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck, _ := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Mismatch")
	}

}
