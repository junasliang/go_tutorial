package main

func main() {
	cards := deck{card(), card()}
	cards = append(cards, "ace of diamonds")

	cards.print()
}

func card() string {
	return "five of diamonds"
}
