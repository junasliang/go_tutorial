package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//alex := person{"Alex", "Lee"}
	alex := person{
		firstName: "Alex",
		lastName:  "Lee",
		contactInfo: contactInfo{
			email:   "alexlee@gmail.com",
			zipCode: 94000,
		},
	}
	alex.updateName("Dicky")
	alex.print()
}

// run to observe pointer features
func (p person) updateName(newFirstname string) {
	p.firstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
