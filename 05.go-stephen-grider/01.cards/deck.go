package main

import (
	"errors"
	"fmt"
)

// Create a new type of deck which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%v of %v", suite, value))
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck, error) {
	if len(d) < handSize {
		return nil, nil, errors.New("not enough cards")
	}

	return d[:handSize], d[handSize:], nil
}
