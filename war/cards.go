// war contains the logic for the deck of cards and gameplay
package war

import "fmt"

var suits = []string{
	"\u2664", // ♤ spade
	"\u2661", // ♡ heart
	"\u2662", // ♢ diamond
	"\u2667", // ♧ club
}

// Aces are high so ranks run from the numbers on the face cards to 11-14 for face cards
var ranks = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

var faces = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// A Card holds the suit and face of the card, but also contains the rank for easy
// comparison of value.
type Card struct {
	Suit string
	Face string
	rank int
}

// String prints out a representation of the card
func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Face, c.Suit)
}

// A Deck is just a slice of cards
type Deck []Card

// NewDeck creates an unshuffled deck
func NewDeck() Deck {
	deck := make(Deck, 0, 52)
	for _, suit := range suits {
		for i, face := range faces {
			deck = append(deck, Card{suit, face, ranks[i]})
		}
	}
	return deck
}
