package main

import (
	"09-deck-of-cards/deck"
	"09-deck-of-cards/utils"
	"fmt"
)

func main() {
	d := deck.New().
		DecksNumber(2).
		Shuffled().
		Filtered(func(filterCard deck.Card) bool {
			return filterCard.Rank >= deck.Ten && filterCard.Suit == deck.Clubs
		})
	utils.PrintArray(d.Cards, 3)
	// sorted := cards.DefaultSorted()
	// utils.PrintArray(sorted.Cards, 3)
	hand1 := d.DealCards(2)
	hand2 := d.DealCards(2)

	fmt.Println("Hand1:", hand1)
	fmt.Println("Hand2:", hand2)
	fmt.Println("Remaining:", d.Cards)
}
