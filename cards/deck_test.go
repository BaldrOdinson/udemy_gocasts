package main

import (
	// "strconv"
	"os"
	"testing"
)

var right_deck_len int = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// right_len := 52
	if len(d) != right_deck_len {
		// t.Errorf("Expected deck length of " + strconv.Itoa(right_len) + ", but got " + strconv.Itoa(len(d)))
		t.Errorf("Expected deck length of %v, but got %v", right_deck_len, len(d))
	}

	first_card := "Ace of Spades"
	if d[0] != first_card {
		t.Errorf("Expected first card of %v, but got %v", first_card, d[0])
	}

	last_card := "King of Clubs"
	if d[len(d)-1] != last_card {
		t.Errorf("Expected last card of %v, but got %v", last_card, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != right_deck_len {
		t.Errorf("Expected %v cards in deck, got %v", right_deck_len, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
