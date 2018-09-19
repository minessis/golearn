// Package declaration...
// The name of the package determines whether you're
// creating an executable or dependency type package...
// Specifically, the word "main" is used to create
// an executable type package...
// If we had used any other name for our package,
// and called `go build` it would not spit out an
// executable file...
package main

import (
	"fmt"
)

// Import other packages that we need...
// The import statement is used to give our package
// access to some code that is written inside another
// package...

// The above is saying "give my package main, all the
// code and all the functionality that is contained
// in package 'fmt'"

// Declare functions, tell Go to do things...
// `func` is short for function...

// One requirement of packages named "main" is that
// a function named "main" which is run automatically,
// must also be defined...

func main() {
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")
	// We use the colon equal when we loop
	// because we're discarding the variables
	// each time we iterate...
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string { // returns a value of type string
	return "Five of Clubs"
}

// `array` is a fixed length list
// `slice` is an array that can grow or shrink
