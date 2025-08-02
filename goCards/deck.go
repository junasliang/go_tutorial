package main

import (
	"fmt"
	"os"
	"strings"
)

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
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// access to saved deck
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// if error, log and exit
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// convert byte slice back to deck
	s := strings.Split(string(bs), ",")
	return deck(s)
}
