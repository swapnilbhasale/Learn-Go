package main

import "fmt"

// Create a new type of deck which is a slice of strings
type deck []string

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
