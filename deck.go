package main

import (
	"fmt"
	"strings"
)

// Create a new type of deck => slice of string
type deck []string

/*
* Receiver function
 */
func (d deck) print() { // (d deck) is the receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
* Function
 */
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

/*
* Multiple return values function
 */
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
