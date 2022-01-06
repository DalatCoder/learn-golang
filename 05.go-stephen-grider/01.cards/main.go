package main

import (
	"fmt"
	"log"
)

func main() {
	cards := newDeck()

	hand, remaningCards, err := deal(cards, 5)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Deck 1")
	hand.print()

	fmt.Println("Deck 2")
	remaningCards.print()
}

