package main

import "fmt"

type englishBot struct{} // concrete type
type spanishBot struct{} // concrete type
type bot interface {     // interface type
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func (englishBot) getGreeting() string {
	return "Hello"
}
