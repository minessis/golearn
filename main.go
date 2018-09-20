package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

// structs are used widely throughout Golang.
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// When we see the star--asterisk--operator in front of a
// type, it means we're working with the pointer to the
// type, and in this case, a pointer to a person
func (p *person) updateName(newFirstName string) {
	// When we see the asterisk operator in front of a pointer
	// it takes the pointer and turns it back into the value
	// the pointer is referencing
	(*p).firstName = newFirstName
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim
	jim.updateName("Jimmy")
	jim.print()
}
