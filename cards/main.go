package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingHand := deal(cards, 5)
	fmt.Println("Hand")
	hand.print()
	fmt.Println("Remaining hand")
	remainingHand.print()
}
