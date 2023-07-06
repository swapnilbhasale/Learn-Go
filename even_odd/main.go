package main

import "fmt"

func main() {
	// create a slice of ints from 0 through 10
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// iterate through the slice. For each element,
	// check to see if the number is even or odd
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
