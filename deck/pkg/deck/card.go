//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	Random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
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

func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, option := range options {
		cards = option(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(card Card) int {
	return int(card.Suit)*int(maxRank) + int(card.Rank)
}

func Shuffle(cards []Card) []Card {
	Random.Shuffle(len(cards)-1, func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Rank: Rank(i), Suit: Joker})
		}
		return cards
	}
}

func Filter(filter func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var filtered []Card
		for _, card := range cards {
			if !filter(card) {
				filtered = append(filtered, card)
			}
		}
		return filtered
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var deck []Card
		for i := 0; i < n; i++ {
			deck = append(deck, cards...)
		}
		return deck
	}
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
