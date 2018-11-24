package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
}

// Reciever to print
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println() // print new line
}

// Create a reciever, that accepts a pointer
func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}
func main() {

	// Assignment with Keys
	p := person{
		firstName: "Alicia",
		lastName:  "Keyz",
		contact: contactInfo{
			email: "@.com",
		},
	}

	// Ordered assignment
	// p := person{"howie", "ho"}
	p.print()

	// All of the following are the same.
	// 1.
	pointToPerson := &p
	pointToPerson.updateFirstName("zia1")
	pointToPerson.print()

	// 2.
	p.updateFirstName("zia2")
	p.print()

	// 3.
	(&p).updateFirstName("zia3")
	p.print()
}
