package main

import "fmt"

//Create a new type of deck
//which is a slice of a string

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValue := []string{"Aces", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
