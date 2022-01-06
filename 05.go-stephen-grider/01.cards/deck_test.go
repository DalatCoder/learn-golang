package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected := 52

	if len(d) != expected {
		t.Errorf("Expected deck length of %v but got %v", expected, len(d))
	}

	fCard := "Ace of Spades"
	if d[0] != fCard {
		t.Errorf("Expected first card of %v, but got %v", fCard, d[0])
	}

	lCard := "King of Clubs"
	if d[len(d)-1] != lCard {
		t.Errorf("Expected last card of %v, but got %v", lCard, d[len(d)-1])
	}
}
