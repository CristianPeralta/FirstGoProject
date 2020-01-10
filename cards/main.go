package main

import "fmt"

func main() {
	cards := []string{"First String", newCard()}
	cards = append(cards, "Last String")
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "New String"
}
