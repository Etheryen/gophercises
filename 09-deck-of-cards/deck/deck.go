package deck

import (
	"math/rand"
	"sort"
)

type Deck struct {
	Cards []Card
}

func New() Deck {
	var deck Deck

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			card := Card{Suit: suit, Rank: rank}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

func (oldDeck Deck) DecksNumber(n int) Deck {
	var d Deck
	for i := 0; i < n; i++ {
		d.Cards = append(d.Cards, oldDeck.Cards...)
	}
	return d
}

func (oldDeck Deck) Filtered(f func(filterCard Card) bool) Deck {
	var d Deck
	for _, c := range oldDeck.Cards {
		if f(c) {
			d.Cards = append(d.Cards, c)
		}
	}
	return d
}

func (oldDeck Deck) JokersAdded(n int) Deck {
	d := copyDeck(oldDeck)
	for i := 0; i < n; i++ {
		d.Cards = append(d.Cards, Card{Suit: Joker, Rank: Rank(i)})
	}
	return d
}

func (oldDeck Deck) Sorted(less func(cards []Card) func(i, j int) bool) Deck {
	d := copyDeck(oldDeck)
	sort.Slice(d.Cards, less(d.Cards))
	return d
}

func (oldDeck Deck) DefaultSorted() Deck {
	d := copyDeck(oldDeck)
	sort.Slice(d.Cards, DefaultLess(d.Cards))
	return d
}

func (oldDeck Deck) Shuffled() Deck {
	d := copyDeck(oldDeck)
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	return d
}

// --- helpers ---

func DefaultLess(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absValue(cards[i]) < absValue(cards[j])
	}
}

func copyDeck(oldDeck Deck) Deck {
	d := Deck{Cards: make([]Card, len(oldDeck.Cards))}
	copy(d.Cards, oldDeck.Cards)
	return d
}

func absValue(c Card) uint8 {
	return uint8(c.Suit)*uint8(maxRank) + uint8(c.Rank)
}
