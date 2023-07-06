package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

/*
 1. (d deck) is called a "receiver" on a function.

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

func (d deck) toString() string {
	// convert deck -> []string -> string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// WriteFile function returns error
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// ReadFile function returns []byte and error
	// if nothing goes wrong with ReadFile, the value of error will be 'nil'.
	// 'nil' in Go means no value
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// error handling
		// option 1: log the error and return a call to newDeck()
		// option 2: log the error and quit the program
		// lets use option 2
		fmt.Println("Error:", err)
		// exit the program using os  package
		os.Exit(1)
	}
	// []byte -> string -> []string -> deck
	// use strings package to separate string on comma -> Split(s string, sep string) []string
	s := strings.Split(string(bs), ",")
	// from above "s" is a []string
	// convert []string to deck type
	return deck(s)
}

func (d deck) shuffle() {
	for index := range d {
		// create a source or seed value
		// time.Now().Unixnano() => int64
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		// generate random number
		// this is a pseudo-random number generator
		// uses a seed value, this is not truly random
		//newPosition := rand.Intn(len(d) - 1)
		// or
		newPosition := r.Intn(len(d) - 1)
		// swap elements
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
