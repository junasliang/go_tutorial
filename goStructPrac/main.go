package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//alex := person{"Alex", "Lee"}
	alex := person{
		firstName: "Alex",
		lastName:  "Lee",
		contact: contactInfo{
			email:   "alexlee@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
