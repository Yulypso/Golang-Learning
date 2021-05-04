package main

func main() {

	cards := newDeck()
	hand, remaining := deal(cards, 4)

	hand.print()
	remaining.print()
	println("--")
	cards.save2File("deck.txt")

	cards2 := newDeckFromFile("deck2.txt")
	cards2.print()

	cards2.shuffle()
	cards2.print()
}
