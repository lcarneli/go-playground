package deck

import (
	"fmt"
	"github.com/lcarneli/go-playground/deck/pkg/deck"
	"math/rand"
	"testing"
)

func TestCardString(t *testing.T) {
	card := deck.Card{Rank: deck.Ace, Suit: deck.Spade}

	expected := "Ace of Spades"
	if fmt.Sprint(card) != expected {
		t.Errorf("got %v, expected %v", card, expected)
	}

	card = deck.Card{Rank: deck.Five, Suit: deck.Diamond}
	expected = "Five of Diamonds"
	if fmt.Sprint(card) != expected {
		t.Errorf("got %v, expected %v", card, expected)
	}

	card = deck.Card{Rank: deck.Ten, Suit: deck.Club}

	expected = "Ten of Clubs"
	if fmt.Sprint(card) != expected {
		t.Errorf("got %v, expected %v", card, expected)
	}

	card = deck.Card{Rank: deck.King, Suit: deck.Heart}
	expected = "King of Hearts"
	if fmt.Sprint(card) != expected {
		t.Errorf("got %v, expected %v", card, expected)
	}

	card = deck.Card{Suit: deck.Joker}
	expected = "Joker"
	if fmt.Sprint(card) != expected {
		t.Errorf("got %v, expected %v", card, expected)
	}
}

func TestNewDeck(t *testing.T) {
	cards := deck.New()
	if len(cards) != 52 {
		t.Errorf("got %v, expected 52", len(cards))
	}
}

func TestNewDeckWithDefaultSort(t *testing.T) {
	cards := deck.New(deck.DefaultSort)
	expected := deck.Card{Rank: deck.Ace, Suit: deck.Spade}
	if cards[0] != expected {
		t.Errorf("got %v, expected %v", cards[0], expected)
	}
}

func TestNewDeckWithSort(t *testing.T) {
	cards := deck.New(deck.Sort(deck.Less))
	expected := deck.Card{Rank: deck.Ace, Suit: deck.Spade}
	if cards[0] != expected {
		t.Errorf("got %v, expected %v", cards[0], expected)
	}
}

func TestNewDeckWithShuffle(t *testing.T) {
	deck.Random = rand.New(rand.NewSource(0))
	cards := deck.New()
	expected := deck.New(deck.Shuffle)
	if cards[44] != expected[0] {
		t.Errorf("got %v, expected %v", cards[44], expected[0])
	}

	if cards[31] != expected[1] {
		t.Errorf("got %v, expected %v", cards[31], expected[1])
	}
}

func TestNewDeckWithJokers(t *testing.T) {
	cards := deck.New(deck.Jokers(2))
	count := 0
	for _, card := range cards {
		if card.Suit == deck.Joker {
			count++
		}
	}
	if count != 2 {
		t.Errorf("got %v, expected 2", count)
	}
}

func TestNewDeckWithFilter(t *testing.T) {
	filter := func(card deck.Card) bool {
		return card.Rank == deck.Ace || card.Rank == deck.King
	}
	cards := deck.New(deck.Filter(filter))
	for _, card := range cards {
		if card.Rank == deck.Ace || card.Rank == deck.King {
			t.Errorf("got %v, expected no Ace and King", card.Rank)
			return
		}
	}
}

func TestNewDeckWithDeck(t *testing.T) {
	cards := deck.New(deck.Deck(2))
	if len(cards) != 104 {
		t.Errorf("got %v, expected 104", len(cards))
	}
}
