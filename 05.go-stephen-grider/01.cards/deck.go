package main

import "fmt"

// Create a new type of deck which is a slice of string
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
