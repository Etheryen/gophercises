package blackjack

import (
	"10-blackjack-game/internal/pkg/deck"
	"fmt"
	"strings"
)

type Hand []deck.Card

type Move string

// TODO: handle doubling down
const (
	Hit   Move = "h"
	Stand Move = "s"
	Split Move = "sp"
)

func Run() {
	cards := deck.New().DecksNumber(3).Shuffled()

	var dealer, player Hand

	dealCardsStart(&cards, []*Hand{&dealer, &player})

	// TODO: check instant dealer blackjack
	// TODO: check instant player blackjack

	// Game loop
	var move Move
	for move != Stand {
		printHands(&dealer, &player, false)
		move = getPlayerMove(false)

		if move == Hit {
			player = append(player, draw(&cards))
		}

		if player.Score() >= 21 {
			break
		}
	}

	if player.Score() > 21 {
		printHands(&dealer, &player, true)
		fmt.Println("Player BUST!!!")
		return
	}

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		dealer = append(dealer, draw(&cards))
	}

	printHands(&dealer, &player, true)

	switch {
	case dealer.Score() > 21:
		fmt.Println("Dealer BUST!!!")
	case player.Score() > dealer.Score():
		fmt.Println("Player WIN!!!")
	case dealer.Score() > player.Score():
		fmt.Println("Dealer WIN!!!")
	case player.Score() == dealer.Score():
		fmt.Println("DRAW!!!")
	}
}

func printHands(d *Hand, p *Hand, isFinal bool) {
	fmt.Println()
	if isFinal {
		fmt.Println("--FINAL HANDS--")
		fmt.Println("Dealer:", d.Score())
		fmt.Println(d.String())
	} else {
		fmt.Println("Dealer:", cardScore((*d)[0]))
		fmt.Println(d.DealerString())
	}
	fmt.Println("Player:", p.Score())
	fmt.Println(p.String())
}

func getPlayerMove(canSplit bool) Move {
	message := "What will you do? (h)it, (s)tand"
	if canSplit {
		message += ", (sp)lit"
	}

	fmt.Println(message)
	var input string
	fmt.Scanln(&input)

	switch strings.ToLower(input) {
	case "h":
		return Hit
	case "s":
		return Stand
	case "sp":
		if canSplit {
			return Split
		}
	}

	fmt.Println("Incorrect move, try again...")
	return getPlayerMove(canSplit)
}

func draw(cards *deck.Deck) deck.Card {
	return cards.DealCards(1)[0]
}

func dealCardsStart(cards *deck.Deck, hands []*Hand) {
	for i := 0; i < 2; i++ {
		for _, h := range hands {
			*h = append(*h, draw(cards))
		}
	}
}

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i, c := range h {
		strs[i] = c.String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", --HIDDEN--"
}

func (h Hand) Score() uint8 {
	score := h.MinScore()

	// Amend ace score
	for aces := h.AcesCount(); score <= 11 && aces > 0; aces-- {
		score += 10
	}

	return score
}

func (h Hand) MinScore() uint8 {
	var minScore uint8

	for _, c := range h {
		minScore += cardScore(c)
	}

	return minScore
}

func cardScore(c deck.Card) uint8 {
	if c.Rank <= deck.Ten {
		return uint8(c.Rank)
	}
	return 10
}

func (h Hand) AcesCount() uint8 {
	var result uint8

	for _, c := range h {
		if c.Rank == deck.Ace {
			result++
		}
	}

	return result
}
