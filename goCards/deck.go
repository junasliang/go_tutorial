package main

import "fmt"

// Create a new type of 'deck'
// slice of strings
type deck []string

// create and return a deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// deal a hand
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//func shuffle()

// save to file

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
