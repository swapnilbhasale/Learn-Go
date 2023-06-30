package main

import "fmt"

// Create a new type of deck which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/* 	1. (d deck) is called a "receiver" on a function.
2. Any variable of type "deck" now gets access to the "print" method.
3. d = The actual copy of the deck we are working with is available
in the function as a variable called 'd'.
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
1. Golang has support to return multiple values like (deck, deck)
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
