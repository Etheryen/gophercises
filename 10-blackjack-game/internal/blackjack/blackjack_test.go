package blackjack

import (
	"10-blackjack-game/internal/pkg/deck"
	"fmt"
	"testing"
)

func TestCardScore(t *testing.T) {
	cases := []struct {
		input deck.Card
		want  int
	}{
		{input: deck.Card{Suit: deck.Clubs, Rank: deck.Queen}, want: 10},
		{input: deck.Card{Suit: deck.Clubs, Rank: deck.King}, want: 10},
		{input: deck.Card{Suit: deck.Clubs, Rank: deck.Jack}, want: 10},
		{input: deck.Card{Suit: deck.Clubs, Rank: deck.Ten}, want: 10},
		{input: deck.Card{Suit: deck.Clubs, Rank: deck.Seven}, want: 7},
		{input: deck.Card{Suit: deck.Clubs, Rank: deck.Two}, want: 2},
	}

	for _, c := range cases {
		t.Run(c.input.String(), func(t *testing.T) {
			got := cardScore(c.input)
			if got != uint8(c.want) {
				t.Errorf("cardScore(%v) = %v, want %v", c.input, got, c.want)
			}
		})
	}
}

func TestHandScore(t *testing.T) {
	cases := []struct {
		input Hand
		want  int
	}{
		{input: Hand{
			{Suit: deck.Clubs, Rank: deck.Queen},
			{Suit: deck.Clubs, Rank: deck.Seven},
		}, want: 17},
		{input: Hand{
			{Suit: deck.Clubs, Rank: deck.King},
			{Suit: deck.Clubs, Rank: deck.Ace},
		}, want: 21},
		{input: Hand{
			{Suit: deck.Clubs, Rank: deck.Eight},
			{Suit: deck.Clubs, Rank: deck.Nine},
			{Suit: deck.Clubs, Rank: deck.Ace},
		}, want: 18},
		{input: Hand{
			{Suit: deck.Hearts, Rank: deck.Ace},
			{Suit: deck.Clubs, Rank: deck.Ace},
		}, want: 12},
		{input: Hand{
			{Suit: deck.Hearts, Rank: deck.Ace},
			{Suit: deck.Clubs, Rank: deck.Ace},
			{Suit: deck.Spades, Rank: deck.Ace},
			{Suit: deck.Diamonds, Rank: deck.Ace},
		}, want: 14},
		{input: Hand{
			{Suit: deck.Hearts, Rank: deck.Ace},
			{Suit: deck.Clubs, Rank: deck.Ace},
			{Suit: deck.Spades, Rank: deck.Ace},
			{Suit: deck.Diamonds, Rank: deck.Ace},
			{Suit: deck.Diamonds, Rank: deck.Ten},
		}, want: 14},
	}

	for i, c := range cases {
		caseName := fmt.Sprintf("case #%v", i+1)
		t.Run(caseName, func(t *testing.T) {
			got := c.input.Score()
			if got != uint8(c.want) {
				t.Errorf("hand.Score(%v) = %v, want %v", c.input, got, c.want)
			}
		})
	}
}
