package main

import "fmt"

var card string

func main() {

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "King of Hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Ace of Spades"
}
