package main

import (
	"fmt"
)

func main() {
	// literal map syntax
	// create a map where:
	// all the keys are strings, and
	// all the values are maps
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
