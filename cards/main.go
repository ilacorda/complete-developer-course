package main


import "fmt"

func main() {
	cards := newDeck()

	// Test the saveToFile() function
	cards.saveToFile("MyCards")

	// Test newDeckFromFile
	newDeck := newDeckFromFile("MyCards")
	fmt.Println("New Deck from File: ", newDeck)

	// Test newDeckFromFile with a non existing filename/file
	errDeck := newDeckFromFile("ADeck")
	fmt.Println("New Deck from File: ", errDeck)

	// Test shuffle() function
	cards.shuffle()
	cards.print()

	// Test deal() function
	hand, remaningCards := deal(cards, 5)

	hand.print()
	remaningCards.print()

	// Example of type conversion
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	// Test the toString() function

	fmt.Println(cards.toString())
}
