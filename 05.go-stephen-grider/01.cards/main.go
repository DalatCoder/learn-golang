package main

import "log"

func main() {
	cards, err := newDeckFromFile("data.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	cards.print()
}

