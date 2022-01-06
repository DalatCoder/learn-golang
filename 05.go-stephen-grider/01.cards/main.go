package main

func main() {
	cards := newDeck()

	err := cards.saveToFile("data.txt")
	if err != nil {
		return 
	}
}

