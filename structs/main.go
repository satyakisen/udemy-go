package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p := person{
		firstName: "Satyaki",
		lastName:  "Sen",
		contactInfo: contactInfo{
			address: "23 A Ghosh Lane",
			zipCode: 700006,
		},
	}

	p.print()
	p.updateName("Satyak")
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newName string) {
	p.firstName = newName
}
