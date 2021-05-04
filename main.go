package main

import "fmt"

func main() {

	cards := newDeck()
	hand, remaining := deal(cards, 4)

	hand.print()
	remaining.print()
	println("--")
	fmt.Println(cards.toString())
}
