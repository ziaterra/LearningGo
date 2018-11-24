package main

func main() {
	// Deck Operations
	cards := newDeck()

	// Print Operations
	// cards.print()
	// fmt.Println(cards.toString())

	// IO Operations
	// cards.saveToFile("my_cards")
	// newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()
}
