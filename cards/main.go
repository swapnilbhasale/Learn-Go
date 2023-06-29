package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// Or can do the following...
	// card := "Ace of Spades"

	//card := newCard()
	//fmt.Println(card)

	// create a slice
	//cards := []string{"Ace of Diamonds", newCard()}
	cards := deck{"Ace of Diamonds", newCard()}
	// append through a slice
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	// iterate over a slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// iterate using the function created in deck.go
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
