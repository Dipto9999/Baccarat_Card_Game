/********************************/
/* Describe Card Deck Behaviour */
/********************************/

package main

import "fmt"

const totalSuites = 4
const totalValues = 14

func initDeck() []string {
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
	playing_deck := []string{}
	for i := 0; i < len(card_suites); i++ {
		for j := 0; j < len(card_values); j++ {
			playing_deck = append(playing_deck, fmt.Sprintf("%s of %s", card_values[j], card_suites[i]))
		}
	}

	return playing_deck
}

func addCard(playing_deck []string, new_card string) []string {
	return append(playing_deck, new_card)
}

func findCard(playing_deck []string, card string) int {
	for i, current := range playing_deck {
		if current == card {
			return i
		}
	}

	return -1
}

func giveCard(playing_deck []string, card string) []string {
	i := findCard(playing_deck, card)

	if i >= 0 {
		playing_deck = append(playing_deck[:i], playing_deck[(i+1):]...)
	}
	return playing_deck
}

func addCards(playing_deck []string, effective_deck []string, number_of_cards int) ([]string, []string) {
	for i := 0; i < number_of_cards; i++ {
		new_card := generateCard(playing_deck)

		playing_deck = giveCard(playing_deck, new_card)
		effective_deck = addCard(effective_deck, new_card)
	}

	return playing_deck, effective_deck
}
