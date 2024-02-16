package deck_test

import (
	"09-deck-of-cards/deck"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	got := len(deck.New().Cards)
	want := 13 * 4

	if got != want {
		t.Errorf("len(deck.New()) = %v, want %v", got, want)
	}
}

func TestShuffle(t *testing.T) {
	original := deck.New()
	shuffled := original.Shuffled()

	if reflect.DeepEqual(original, shuffled) {
		t.Error("Original deck is the same as shuffled")
	}
}

func TestDefaultSort(t *testing.T) {
	original := deck.New()
	sorted := original.Shuffled().DefaultSorted()

	if !reflect.DeepEqual(original, sorted) {
		t.Error("Original deck is different than sorted")
	}
}

func TestSort(t *testing.T) {
	original := deck.New()
	sorted := original.Shuffled().Sorted(deck.DefaultLess)

	if !reflect.DeepEqual(original, sorted) {
		t.Error("Original deck is different than sorted")
	}
}

func TestJokers(t *testing.T) {
	n := 3
	d := deck.New().Filtered(func(filterCard deck.Card) bool {
		return filterCard.Rank == deck.Four
	})

	d = d.JokersAdded(n)

	for i := len(d.Cards) - n; i < len(d.Cards); i++ {
		if d.Cards[i].Suit != deck.Joker {
			t.Errorf(
				"cards[i] = %v, want %v",
				d.Cards[i].String(),
				deck.Joker.String(),
			)
		}
	}
}

func TestFilter(t *testing.T) {
	d := deck.New().Filtered(func(c deck.Card) bool {
		return c.Rank != deck.Two && c.Rank != deck.Three &&
			c.Suit != deck.Diamonds
	})

	if len(d.Cards) == 0 {
		t.Error("expected cards not to be empty")
	}

	for _, c := range d.Cards {
		if c.Rank == deck.Two || c.Rank == deck.Three ||
			c.Suit == deck.Diamonds {
			t.Error("expected all twos, threes and diamonds to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	n := 4
	d := deck.New().DecksNumber(n)

	if len(d.Cards) != 13*4*n {
		t.Errorf("len(cards) = %v, want %v", len(d.Cards), 13*4*n)
	}
}

func TestDealCards(t *testing.T) {
	d := deck.New().
		DecksNumber(3).
		Filtered(func(filterCard deck.Card) bool {
			return filterCard.Rank > deck.Ten && filterCard.Suit == deck.Clubs
		})

	hand1 := d.DealCards(2)
	hand2 := d.DealCards(2)

	want1 := []deck.Card{
		{Suit: deck.Clubs, Rank: deck.Jack},
		{Suit: deck.Clubs, Rank: deck.Queen},
	}
	want2 := []deck.Card{
		{Suit: deck.Clubs, Rank: deck.King},
		{Suit: deck.Clubs, Rank: deck.Jack},
	}
	wantRemaining := []deck.Card{
		{Suit: deck.Clubs, Rank: deck.Queen},
		{Suit: deck.Clubs, Rank: deck.King},
		{Suit: deck.Clubs, Rank: deck.Jack},
		{Suit: deck.Clubs, Rank: deck.Queen},
		{Suit: deck.Clubs, Rank: deck.King},
	}

	if !reflect.DeepEqual(hand1, want1) {
		t.Errorf("hand1 = %v, want %v", hand1, want1)
	}

	if !reflect.DeepEqual(hand2, want2) {
		t.Errorf("hand2 = %v, want %v", hand2, want2)
	}

	if !reflect.DeepEqual(d.Cards, wantRemaining) {
		t.Errorf("d.Cards = %v, want %v", d.Cards, wantRemaining)
	}
}
