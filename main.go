package main

import "fmt"

func main() {
	var card string = "Ace of Spades"              //Used when variable is initialized
	card1 := "Alternate way of decalring variable" //Used when variable is initialized
	fmt.Println(card)
	fmt.Println(card1)
	card2 := newCard()
	fmt.Println(card2)
	card3 := []string{newCard(), "Aces"}
	fmt.Println(card3)
	card3 = append(card3, "Jokers")

	for i, c := range card3 {
		fmt.Println(i, c)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
