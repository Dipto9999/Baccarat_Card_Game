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
	// Get a new source.
	r := rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)

	// Generate a random nmber within specified range.
	return r.Intn(max)
}

// Create a New Type Deck (Slice of Strings)
type deck []string

/*
initDeck creates a new deck of cards, represented as a slice of strings. The deck
is made of 52 cards, with 4 suits (Clubs, Diamonds, Hearts, Spades) and 14 values
(Ace, 1-10, Jack, Queen, King) for each suit. The cards are ordered first by suit,
and then by value within each suit.
*/
func initDeck() deck {
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

	return game_deck
}

/*
newDeck creates a new deck of cards, represented as a slice of strings. The deck
contains 52 cards, with 4 suits (Clubs, Diamonds, Hearts, Spades) and 14 values
(Ace, 1-10, Jack, Queen, King) for each suit. The cards are ordered first by suit,
and then by value within each suit.

The function initializes arrays for suits and values, and then iterates through them
to create the deck. The deck is then shuffled using the `shuffle` method, printed to
the console using the `print` method, and saved to a file named "game_deck" using
the `saveToFile` method.

The function returns the created deck as a `deck` type.
*/
func newDeck() deck {
	// Init Deck
	game_deck := initDeck()

	game_deck.shuffle(5)
	game_deck.print()
	game_deck.saveToFile("game_deck.txt")

	return game_deck
}

/*
toString returns a single string representation of the entire deck of cards, or
a string representation of each individual card, separated by newlines. If the
`individual` argument is supplied and non-empty, the function will return a string
representation of each card on its own line. Otherwise, the function will return a
single string representing the entire deck, with each card separated by a newline.

The function iterates through the `deck` type and appends each card to a single
string, using the `fmt.Sprintf` method to format each card string as "%s\n". The
resulting string is then returned.

If the `individual` argument is supplied, the resulting string will contain one
card per line, with each line terminated by a newline character ("\n"). Otherwise,
the resulting string will contain the entire deck as a single string, with each
card separated by a newline character.
*/
func (d deck) toString(individual ...string) string {
	converted := ""
	for i := 0; i < len(d); i++ {
		if individual != nil {
			converted += fmt.Sprint(d[i])
		} else {
			converted += fmt.Sprint(d[i])
		}

		if i < len(d)-1 {
			converted += "\n"
		}
	}

	return converted
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

/*
print outputs the contents of the deck of cards to the console, either as a
single list or as individual cards, depending on the `individual` argument.

If the `individual` argument is supplied and non-empty, the function will print
each card in the deck on its own line, with a prefix string indicating the source
of the cards (i.e. the name of the deck). Each line will be formatted as:
"{prefix} Card #{index} : {card}"

If the `individual` argument is not supplied, the function will print the entire
deck as a single list, with each card on its own line. Each line will be formatted
as: "Card #{index} : {card}"

The function iterates through the `deck` type and prints each card to the console,
using the `fmt.Printf` method to format each card string according to the format
described above. The output is then terminated with a newline character ("\n").
*/
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

/*
getCard adds a new card to the end of the deck of cards.

The `card` argument is a string representing the new card to be added to the deck.

The function modifies the `deck` type by appending the new card to the end of the
slice, using the `append` method. The modification is made to the original `deck`
value, which is passed as a pointer (`*deck`).
*/
func (d *deck) getCard(card string) {
	*d = append(*d, card)
}

/*
giveCard removes a random card from the deck of cards and returns it as a string.

The function generates a random index `i` within the range of the deck's length,
using the `getNumber` function, which is assumed to return a random integer between
0 (inclusive) and `n` (exclusive), where `n` is the length of the deck. The card at
index `i` is then retrieved from the slice and stored in the `card` variable.

The function modifies the `deck` type by removing the card at index `i` from the slice,
using slice indexing and the `append` method. The modification is made to the original
`deck` value, which is passed as a pointer (`*deck`).

Finally, the function returns the `card` variable, which contains the string value of
the card that was removed from the deck.
*/
func (d *deck) giveCard() string {
	i := getNumber(len(*d))
	card := (*d)[i]

	*d = append((*d)[:i], (*d)[(i+1):]...)
	return card
}

/*
shuffle shuffles the order of cards in the deck using a random swapping algorithm.

The `iter` argument determines the number of times the shuffling process is repeated.
If `iter` is less than 0, it is set to 1 to ensure at least one iteration is performed.

The function performs the shuffling process by iterating `iter` times. In each iteration,
it iterates over each card in the deck using the range loop and swaps the current card with
a randomly selected card. The random index `j` is generated using the `getNumber` function,
which is assumed to return a random integer between 0 (inclusive) and `n` (exclusive), where
`n` is the length of the deck. The swap is performed using the multiple assignment syntax
`(*d)[i], (*d)[j] = (*d)[j], (*d)[i]`.

The shuffling process modifies the order of cards in the original `deck` value, which is
passed as a pointer (`*deck`).
*/
func (d *deck) shuffle(iter int) {
	if iter < 0 {
		iter = 1
	}

	for rep := 0; rep < iter; rep++ {
		for i := range *d {
			var j = getNumber(len(*d))
			(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
		}
	}
}

/*
deal deals a specified number of cards from one deck to another.

The `tx` argument represents the deck from which the cards are being dealt, and `rx`
represents the deck to which the cards are being dealt. The `num_cards` argument determines
the number of cards to be dealt. If `num_cards` is less than 0, it is set to 1 to ensure
at least one card is dealt.

The function performs the dealing process by iterating `num_cards` times. In each iteration,
it retrieves a card from the `tx` deck by using the `giveCard` method, and adds that card
to the `rx` deck by using the `getCard` method. This process transfers a card from `tx` to `rx`.

Additionally, after each card is dealt, the `tx` deck is saved to a file named "game_deck"
using the `saveToFile` method. This allows the deck to be updated and persisted after each
deal operation.

The dealing process modifies the contents of both the `tx` and `rx` decks, which are passed
as pointers (`*deck`).
*/
func (tx *deck) deal(rx *deck, num_cards int) {
	if num_cards < 0 {
		num_cards = 1
	}

	for i := 0; i < num_cards; i++ {
		rx.getCard(tx.giveCard())

		(*tx).saveToFile("game_deck.txt")
	}
}
