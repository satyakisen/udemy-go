package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 6 {
		t.Errorf("length of card expected is 6 but found %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected element to be Ace of Spades but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Three of Diamonds" {
		t.Errorf("Expected element to be Three of Diamonds but got %v", cards[len(cards)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	cards := newDeck()

	cards.saveToFile("_decktesting")

	loadedCards := newDeckFromFile("_decktesting")

	if len(loadedCards) != 6 {
		t.Errorf("length of card expected is 6 but found %v", len(loadedCards))
	}

	os.Remove("_decktesting")
}
