package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

/*
 * all type and custom type who
 * has a func named getGreeting and
 * has a return value string
 * are now a bot.
 */

type bot interface { // Can't create a interface type value but only concrete type
	getGreeting() string
	/*
	 * add other required func to become a custom type bot
	 */
}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
