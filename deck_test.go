/*************/
/* Test File */
/*************/

package main

import (
	"testing"
)

func TestInitDeck(t *testing.T) {
	d := initDeck()

	// Check that d is Made of 56 Cards.
	if len(d) != 56 {
		t.Errorf("d Length of %d. Expected 56", len(d))
	}

	// Check that First Card is Ace of Clubs.
	if d[0] != "Ace of Clubs" {
		t.Errorf("First Card is %s. Expected Ace of Clubs", d[0])
	}

	// Check that Last Card is 10 of Spades.
	if d[len(d)-1] != "10 of Spades" {
		t.Errorf("Last Card is %s. Expected 10 of Spades", d[len(d)-1])
	}
}

func TestNewDeck(t *testing.T) {
	d_returned := newDeck()
	d_file := readFromFile("game_deck.txt")

	// Check that the d is made of 56 cards.
	if len(d_returned) != 56 {
		t.Errorf("d Length of %d. Expected 56", len(d_returned))
	}

	// Check that d Written to File is Made of 56 Cards.
	if len(d_file) != len(d_returned) {
		t.Errorf("d_file Length of %d. Expected 56", len(d_file))
	}

	// Check that Both Decks Are the Same.
	for i := 0; i < len(d_returned); i++ {
		if d_file[i] != d_returned[i] {
			t.Errorf("Card %d in File Deck is %s.", i, d_file[i])
			t.Errorf("Card %d in Returned Deck is %s.", i, d_returned[i])
		}
	}
}

func TestShuffle(t *testing.T) {
	d_unshuffled := initDeck()

	// Copy Unshuffled Deck to Shuffled Deck.
	d_shuffled := deck{}
	for _, card := range d_unshuffled {
		d_shuffled = append(d_shuffled, card)
	}

	// Shuffle Deck And Check That It Is Different From Unshuffled Deck
	shuffle_count := 0
	shuffled := false
	for shuffled == false {
		d_shuffled.shuffle(1) // Shuffle d_shuffled.
		for _, card := range d_unshuffled {
			if d_shuffled[0] != card {
				shuffled = true
			}
		}

		// If Deck Is Not Shuffled After 1000 Shuffles, Throw Error.
		if shuffle_count > 1000 {
			t.Errorf("Deck is Not Shuffled After %d Shuffles", shuffle_count)
		} else {
			shuffle_count++ // Increment Shuffle Count.
		}
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()

	my_cards := deck{}

	d.deal(&my_cards, 5)

	if len(my_cards) != 5 {
		t.Errorf("My Cards Length of %d. Expected 5", len(my_cards))
	}

	if len(d) != 51 {
		t.Errorf("d Length of %d. Expected 51", len(d))
	}

	// Ensure That No Cards Are In Both Decks.
	found_card := ""
	for _, card_i := range d {
		for _, card_j := range my_cards {
			if card_i == card_j {
				found_card += card_i
			}
		}
	}

	if found_card != "" {
		t.Errorf("Found Card %s In Both Decks", found_card)
	}
}
