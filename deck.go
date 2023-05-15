/********************************/
/* Describe Card Deck Behaviour */
/********************************/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	totalValues = 14
	totalSuits  = 4
)

var values [totalValues]string
var suits [totalSuits]string

// Card represents a playing card.
type card struct {
	value string
	suit  string
}

// Deck represents a deck of cards.
type deck struct {
	owner string
	cards []card
}

/*
init initializes the values and suits arrays.
*/
func init() {
	values = [totalValues]string{
		"Ace", "Jack", "Queen", "King",
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
	}

	suits = [totalSuits]string{
		"Clubs", "Diamonds", "Hearts", "Spades",
	}
}

/*
getNumber generates a random number within the specified range.

Parameters:
- max: An integer representing the upper bound (exclusive) of the range.

Returns:
- random number within the range [0, max) as an integer.
*/
func getNumber(max int) int {
	// Get a new source.
	r := rand.New(
		rand.NewSource(
			time.Now().UnixNano(),
		),
	)

	// Generate a random number within a specified range.
	return r.Intn(max)
}

/*
initDeck creates a new deck of cards by iterating through the suits and values.
It appends each combination of suit and value as a card to the deck.

Parameters:
- individual (optional): variadic parameter representing the owner of the deck.

Returns:
-deck struct representing a new deck of cards with all possible combinations of suits and values.
*/
func initDeck(individual ...string) deck {
	game_deck := deck{}
	// Iterate through suits and values to create a deck of cards.
	for _, suit := range suits {
		for _, value := range values {
			game_deck.cards = append(
				game_deck.cards,
				card{
					value: value,
					suit:  suit,
				},
			)
		}
	}
	// Identify the owner of the deck.
	if individual != nil {
		game_deck.owner = individual[0]
	}

	return game_deck
}

/*
newDeck creates a new deck of cards, shuffles it, prints it, and saves it to a file.

Parameters:
- times_shuffled: An integer specifying the number of times the deck should be shuffled.
- individual (optional): A variadic parameter representing the owner of the deck.

Returns:
- deck struct representing a new deck of cards with all possible combinations of suits and values.
*/
func newDeck(times_shuffled int, individual ...string) deck {
	// Init deck of cards.
	game_deck := initDeck(individual...)

	game_deck.shuffle(times_shuffled)
	game_deck.print()

	// Save the deck of cards to a file.
	game_deck.saveToFile(
		strings.ToLower(strings.ReplaceAll(game_deck.owner, " ", "_")) + ".txt",
	)

	return game_deck
}

/*
toString takes a deck of cards and iterates through each card to
produce a string representing the entire deck of cards.

Receiver:
- d: deck to convert to a string.

Returns:
- deck struct as a string.
*/
func (d deck) toString() string {
	converted := ""

	// Iterate through the deck of cards and convert each card to a string.
	for i := 0; i < len(d.cards); i++ {
		converted += fmt.Sprintf("%s of %s", d.cards[i].value, d.cards[i].suit)

		// Add a newline character to the end of each card, except the last.
		if i < (len(d.cards) - 1) {
			converted += "\n"
		}
	}

	return converted
}

/*
toByteSlice converts the deck of cards to a byte slice.

Receiver:
- d: deck to convert to a byte slice.

Returns:
- deck struct as a byte slice.
*/
func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

/*
print prints the deck of cards.

Receiver:
- d: deck struct to print.
*/
func (d deck) print() {
	// Iterate through the deck of cards and convert each card to a string.
	converted := ""
	for i := 0; i < len(d.cards); i++ {
		converted += fmt.Sprintf("%s Card #%d : %s of %s\n",
			d.owner, (i + 1), d.cards[i].value, d.cards[i].suit,
		)
	}

	converted += "\n"
	fmt.Print(converted)
}

/*
getCard appends a card to the deck of cards.

Receiver:
- d: deck to which the card will be appended.

Parameters:
- c: card to append to the deck.
*/
func (d *deck) getCard(c card) {
	(*d).cards = append((*d).cards, c)
}

/*
giveCard retrieves a random card from the deck and removes it.

Receiver:
- d: deck from which the card will be retrieved.

Returns:
- card struct representing the card retrieved from the deck.
*/
func (d *deck) giveCard() card {
	// Get a random card from the deck.
	i := getNumber(len((*d).cards))
	c := (*d).cards[i]

	// Remove the card from the deck.
	(*d).cards = append((*d).cards[:i], (*d).cards[(i+1):]...)
	return c
}

/*
shuffle iterates through the deck of cards a specified number of times
and swaps each card with a random card.

Receiver:
- d: deck to shuffle.

Parameters:
- iter: An integer representing the number of times to iterate through the deck.
*/
func (d *deck) shuffle(iter int) {
	if iter < 0 {
		iter = 1
	}

	for rep := 0; rep < iter; rep++ {
		for i := range (*d).cards {
			// Get a random card from the deck and swap it with the current card.
			var j = getNumber(len((*d).cards))
			((*d).cards)[i], ((*d).cards)[j] = ((*d).cards)[j], ((*d).cards)[i]
		}
	}
}

/*
deal takes a specified number of randomly selected cards from the deck
and gives them to another deck. It then updates the file for the deck
from which the cards were taken.

Receiver:
- tx: deck from which cards will be taken.

Parameters:
- rx: deck to which cards will be given.
- num_cards: An integer representing the number of cards to take from the deck.
*/
func (tx *deck) deal(rx *deck, num_cards int) {
	if num_cards < 0 {
		num_cards = 1
	}

	for i := 0; i < num_cards; i++ {
		rx.getCard(tx.giveCard())
	}

	// Save the deck of cards to a file.
	tx.saveToFile(
		strings.ToLower(strings.ReplaceAll(tx.owner, " ", "_")) + ".txt",
	)
}
