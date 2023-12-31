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
	//cards := deck{"Ace of Diamonds", newCard()}
	cards := newDeck()

	// append through a slice
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	// iterate over a slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// iterate using the function created in deck.go
	cards.print()

	// deal function returns 2 values
	/*
		We create 2 new references that point at subsections of the 'cards'
		slice. We did not directly modify the slice that 'cards' is pointing at.
	*/
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	// convert deck to string
	// this function will return comma separated deck
	// defined in deck.go
	fmt.Println(cards.toString())

	// save cards to file
	cards.saveToFile("my_cards")

	// load deck from file
	cards_from_file := newDeckFromFile("my_cards")
	cards_from_file.print()

	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
