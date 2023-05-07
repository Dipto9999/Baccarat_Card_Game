/********************************/
/* Describe Card Deck Behaviour */
/********************************/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const totalSuites = 4
const totalValues = 14

// Create a New Type Deck (Slice of Strings)
type deck []string

func newDeck() deck {
	// Init Array of Strings for Card Suites.
	card_suites := [totalSuites]string{
		"Clubs", "Diamonds", "Hearts", "Spades",
	}

	// Init Array of Strings for Card Values
	card_values := [totalValues]string{
		"Ace", "Jack", "Queen", "King",
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
	}

	// Iterate Through Arrays and Init Deck.
	game_deck := deck{}
	for i := 0; i < len(card_suites); i++ {
		for j := 0; j < len(card_values); j++ {
			game_deck = append(game_deck, fmt.Sprintf("%s of %s", card_values[j], card_suites[i]))
		}
	}

	game_deck.print()
	return game_deck
}

func getNumber(max_number int) int {
	// Seed the Random Number Generator.
	rand.Seed(time.Now().UnixNano())

	// Generate a Random Number Within Specified Range.
	return rand.Intn(max_number)
}

func (cards deck) print(individual ...string) {
	for i, card := range cards {
		if individual != nil {
			fmt.Printf("%s Card #%d : %s\n", individual[0], (i + 1), card)
		} else {
			fmt.Printf("Card #%d : %s\n", (i + 1), card)
		}
	}
	fmt.Printf("\n")
}

func (cards *deck) _addCard(new_card string) {
	*cards = append(*cards, new_card)
}

func (cards *deck) _giveCard() string {
	i := getNumber(len(*cards))
	card := (*cards)[i]

	*cards = append((*cards)[:i], (*cards)[(i+1):]...)
	return card
}

func (rx *deck) addCards(tx *deck, num_cards int) {
	if num_cards < 0 {
		num_cards = 1
	}

	for i := 0; i < num_cards; i++ {
		rx._addCard(tx._giveCard())
	}
}
