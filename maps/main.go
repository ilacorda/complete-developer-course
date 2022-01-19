package main

import "fmt"

func main() {
	// first way to declare a map
	var someColors map[string]string
	fmt.Println("someColors ", someColors)

	// second way to declare a map
	anotherColors := make(map[string]string)
	fmt.Println("anotherColors ", anotherColors)

	anotherColors["white"] = "#fffff"

	fmt.Println("anotherColors with White ", anotherColors)

	delete(anotherColors, "White")

	fmt.Println("anotherColors now empty ", anotherColors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	fmt.Println("colors ", colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)

	}

}
