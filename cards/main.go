package main

import "fmt"

func main() {
	cards := deck{"First String", newCard()}
	cards = append(cards, "Last String")
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "New String"
}
