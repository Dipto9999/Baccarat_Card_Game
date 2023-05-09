/********************************/
/* Describe Card Deck Behaviour */
/********************************/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const totalSuits = 4
const totalValues = 14

func getNumber(max int) int {
	// Seed the Random Number Generator.
	rand.Seed(time.Now().UnixNano())

	// Generate a Random Number Within Specified Range.
	return rand.Intn(max)
}

// Create a New Type Deck (Slice of Strings)
type deck []string

func newDeck() deck {
	// Init Array of Strings for Card Suits.
	suits := [totalSuits]string{
		"Clubs", "Diamonds", "Hearts", "Spades",
	}

	// Init Array of Strings for Card Values
	values := [totalValues]string{
		"Ace", "Jack", "Queen", "King",
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
	}

	// Iterate Through Arrays and Init Deck.
	game_deck := deck{}
	for _, suit := range suits {
		for _, value := range values {
			game_deck = append(game_deck, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	game_deck.shuffle(3)
	game_deck.print()
	game_deck.saveToFile("game_deck")

	return game_deck
}

func (d deck) toString(individual ...string) string {
	converted := ""
	for _, card := range d {
		if individual != nil {
			converted += fmt.Sprintf("%s\n", card)
		} else {
			converted += fmt.Sprintf("%s\n", card)
		}
	}

	return converted
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) print(individual ...string) {
	for i, card := range d {
		if individual != nil {
			fmt.Printf("%s Card #%d : %s\n", individual[0], (i + 1), card)
		} else {
			fmt.Printf("Card #%d : %s\n", (i + 1), card)
		}
	}

	fmt.Printf("\n")
}

func (d *deck) getCard(card string) {
	*d = append(*d, card)
}

func (d *deck) giveCard() string {
	i := getNumber(len(*d))
	card := (*d)[i]

	*d = append((*d)[:i], (*d)[(i+1):]...)
	return card
}

func (d *deck) shuffle(iter int) {
	if iter < 0 {
		iter = 1
	}

	for rep := 0; rep < iter; rep++ {
		for i := 0; i < len(*d); i++ {
			var curr = (*d)[i]
			(*d)[i] = (*d)[getNumber(len(*d))]
			(*d)[getNumber(len(*d))] = curr
		}
	}
}

func (tx *deck) deal(rx *deck, num_cards int) {
	if num_cards < 0 {
		num_cards = 1
	}

	for i := 0; i < num_cards; i++ {
		rx.getCard(tx.giveCard())

		(*tx).saveToFile("game_deck")
	}
}
