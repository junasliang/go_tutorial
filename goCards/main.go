package main

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	// hand, remainDeck := deal(cards, 5)
	// hand.print()
	// remainDeck.print()
	// cards.saveToFile("my_cards")
	cards.print()
}

func card() string {
	return "five of diamonds"
}
