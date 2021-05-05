package main

import "fmt"

func main() {

	colors1 := map[string]string{ //map[key]value
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	var colors2 map[string]string

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"

	fmt.Println(colors1)
	fmt.Println(colors2)
	fmt.Println(colors3)

	delete(colors3, "white")
	fmt.Println(colors3)

	printMap(colors1)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
