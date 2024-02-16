package deck_test

import (
	"09-deck-of-cards/deck"
	"testing"
)

func TestCardString(t *testing.T) {
	card := deck.Card{Suit: deck.Hearts, Rank: deck.Nine}
	expectedString := "Nine of Hearts"

	if card.String() != expectedString {
		t.Errorf(
			"Card{deck.Hearts, deck.Nine}.String() = %v, want %v",
			card.String(),
			expectedString,
		)
	}

	card = deck.Card{Suit: deck.Joker}
	expectedString = "Joker"

	if card.String() != expectedString {
		t.Errorf(
			"Card{deck.Joker}.String() = %v, want %v",
			card.String(),
			expectedString,
		)
	}
}
