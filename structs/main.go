package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var john person
	john.firstName = "John"
	john.lastName = "Vann"
	fmt.Println(john)
	fmt.Printf("%+v\n", john)
}
