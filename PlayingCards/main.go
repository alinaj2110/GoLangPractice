package main

import (
	"fmt"
	"os"
)

var card string

func main() {

	cards := newDeck()

	fmt.Println(cards.toString())
	cards.saveToFile("new_cards.txt")
	newdeck, err := newDeckFromFile("new_cards.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	newdeck.shuffleDeck()
	newdeck.print()
}
