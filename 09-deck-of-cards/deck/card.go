//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker
)

var suits = [...]Suit{
	Spades,
	Diamonds,
	Clubs,
	Hearts,
}

type Rank uint8

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%v of %v", c.Rank.String(), c.Suit.String())
}
