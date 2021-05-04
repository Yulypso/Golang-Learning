package main

import "fmt"

func main() {

	cards := newDeck()
	hand, remaining := deal(cards, 4)

	hand.print()
	remaining.print()

	d := "hi there"
	fmt.Println([]byte(d)) // to cast into byte slice
}
