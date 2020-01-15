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
	john := person{
		firstName: "John",
		lastName:  "Vann",
		contact: contactInfo{
			email:   "john@gmail.com",
			zipCode: 95000,
		},
	}

	fmt.Printf("%+v\n", john)
}
