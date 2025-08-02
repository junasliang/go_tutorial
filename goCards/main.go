package main

func main() {
	cards := newDeck()
	// hand, remainDeck := deal(cards, 5)
	// hand.print()
	// remainDeck.print()
	cards.saveToFile("my_cards")
	// fmt.Println(cards)
}

func card() string {
	return "five of diamonds"
}
