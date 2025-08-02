package main

func main() {
	cards := newDeck()
	hand, remainDeck := deal(cards, 5)
	hand.print()
	remainDeck.print()
}

func card() string {
	return "five of diamonds"
}
