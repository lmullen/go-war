package war

import "fmt"

var suits = []string{
	"\u2664", // ♤ spade
	"\u2661", // ♡ heart
	"\u2662", // ♢ diamond
	"\u2667", // ♧ club
}

// Aces high
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

var Deck []Card

func init() {

	// Create the deck of cards
	var card Card
	for _, suit := range suits {
		for i, face := range faces {
			card = Card{suit, face, ranks[i]}
			Deck = append(Deck, card)
		}
	}
}
