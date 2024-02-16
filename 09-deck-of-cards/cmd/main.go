package main

import (
	"09-deck-of-cards/deck"
	"09-deck-of-cards/utils"
)

func main() {
	cards := deck.New().
		DecksNumber(2).
		Shuffled().
		Filtered(func(filterCard deck.Card) bool {
			return filterCard.Rank == deck.King
		})
	utils.PrintArray(cards.Cards, 3)
	sorted := cards.DefaultSorted()
	utils.PrintArray(sorted.Cards, 3)
}
