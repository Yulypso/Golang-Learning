package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	/* create chanenl */
	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	/* wait data and print it out through channel */
	/*for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}*/

	/* better implementation
	 * whenever we receive a channel value, we spawn up a go routine
	 */

	/* infinite receiver loop 1*/
	/*for {
		go checkLink(<-c, c)
	}*/

	/* infinite receiver loop 2*/
	/*for l := range c { // infinite loop which waits for a value in c, then assign it to l and finally call checkLink with it

		// pause the code for 5 seconde but it's definitively not appropriate
		//  because the main go routine won't be ready to receive new channel value.

		//time.Sleep(5 * time.Second) // not the best solution
		go checkLink(l, c)
	}*/

	/* infinite receiver loop 3 with function literal == lambda function / anonymous function */
	for l := range c { // l is changing in the main scope
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // func argument l that we pass to the func literal
	}
}

func checkLink(link string, c chan string) {
	//time.Sleep(5 * time.Second) // not the best solution too because we would not check directly the link.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // send data through channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link // send data through channel
}
