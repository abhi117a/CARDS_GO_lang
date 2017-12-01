package main

import "fmt"

func main() {
	var card string = "Ace of Spades"              //Used when variable is initialized
	card1 := "Alternate way of decalring variable" //Used when variable is initialized
	fmt.Println(card)
	fmt.Println(card1)
	card2 := newCard()
	fmt.Println(card2)
	card3 := deck{newCard(), "Aces"}
	fmt.Println(card3)
	card3 = append(card3, "Jokers")
	card3.print()

	card4 := newDeck()
	fmt.Println(card4)

}

func newCard() string {
	return "Five of Diamonds"
}
