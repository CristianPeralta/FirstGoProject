package main

func main() {
	cards := deck{"First String", newCard()}
	cards = append(cards, "Last String")
	cards.print()
}

func newCard() string {
	return "New String"
}
