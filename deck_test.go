/***************************/
/* Test Deck Functionality */
/***************************/

package main

import "testing"

func TestInitDeck(t *testing.T) {
	d := initDeck("Deck Test")

	// Check that d is made of 56 cards.
	if len(d.cards) != 56 {
		t.Errorf("Deck Length of %d. Expected 56", len(d.cards))
	}

	// Check that first card is Ace of Clubs.
	if (d.cards[0].value != "Ace") || (d.cards[0].suit != "Clubs") {
		t.Errorf("First Card is %s. Expected Ace of Clubs", d.cards[0].value+" of "+d.cards[0].suit)
	}

	// Check that last card is 10 of Spades.
	if (d.cards[len(d.cards)-1].value != "10") || (d.cards[len(d.cards)-1].suit != "Spades") {
		t.Errorf("Last Card is %s. Expected 10 of Spades", d.cards[len(d.cards)-1].value+" of "+d.cards[len(d.cards)-1].suit)
	}

	// Check that owner is Deck Test.
	if d.owner != "Deck Test" {
		t.Errorf("Owner is %s. Expected Deck Test", d.owner)
	}

}

func TestNewDeck(t *testing.T) {
	d_returned := newDeck(5, "Deck Test")
	d_file := readFromFile("deck_test.txt")

	// Check that the new deck is made of 56 cards.
	if len(d_returned.cards) != 56 {
		t.Errorf("Returned Deck Length of %d. Expected 56", len(d_returned.cards))
	}

	// Check that the deck written to file is made of 56 cards.
	if len(d_file.cards) != len(d_returned.cards) {
		t.Errorf("File Deck Length of %d. Expected 56", len(d_file.cards))
	}

	// Check that both decks are the same.
	for i := 0; i < len(d_file.cards); i++ {
		if (d_file.cards[i].value != d_returned.cards[i].value) || (d_file.cards[i].suit != d_returned.cards[i].suit) {
			t.Errorf("Card %d in Returned Deck is %s.", i, d_returned.cards[i].value+" of "+d_returned.cards[i].suit)
			t.Errorf("Card %d in File Deck is %s.", i, d_file.cards[i].value+" of "+d_file.cards[i].suit)
		}
	}
}

func TestShuffle(t *testing.T) {
	d := initDeck()

	// Copy unshuffled deck to shuffled deck.
	d_shuffled := deck{}
	for _, card := range d.cards {
		d_shuffled.cards = append(d_shuffled.cards, card)
	}

	// Check that the cards have been shuffled.
	shuffled, shuffle_count := false, 0
	for shuffled == false {
		d_shuffled.shuffle(1) // Shuffle deck.

		// Break if deck is shuffled.
		for i := 0; i < len(d.cards); i++ {
			if (d_shuffled.cards[i].value != d.cards[i].value) || (d_shuffled.cards[i].suit != d.cards[i].suit) {
				shuffled = true
				break
			}
		}

		// If Deck Is Not Shuffled After 1000 Shuffles, Throw Error.
		if shuffle_count > 1000 {
			t.Errorf("Deck is Not Shuffled After %d Shuffles", shuffle_count)
		} else {
			shuffle_count++ // Increment shuffle count.
		}
	}
}

func TestDeal(t *testing.T) {
	d_game := newDeck(5, "Deck Test")
	d_player := deck{}

	d_game.deal(&d_player, 5)

	// Check that d_player is made of 5 cards.
	if len(d_player.cards) != 5 {
		t.Errorf("Player Cards Length of %d. Expected 5", len(d_player.cards))
	}

	// Check that d_game is made of 51 cards.
	if len(d_game.cards) != 51 {
		t.Errorf("Game Deck Length of %d. Expected 51", len(d_game.cards))
	}

	// Ensure that no cards are in both decks.
	c_found := []card{}
	for _, c_game := range d_game.cards {
		for _, c_player := range d_player.cards {
			// If card is in both decks, add to c_found.
			if c_game == c_player {
				c_found = append(c_found, c_game)
				t.Errorf("Found Card %s In Both Decks", c_game)
			}
		}
	}

	// Check that no cards are in both decks.
	if len(c_found) > 0 {
		t.Errorf("Found %d Cards In Both Decks", len(c_found))
	}

	// Check that the deck written to file is made of 51 cards.
	d_file := readFromFile("deck_test.txt")
	if len(d_file.cards) != len(d_game.cards) {
		t.Errorf("File Deck Length of %d. Expected 51", len(d_file.cards))
	}

	// Check that both decks are the same.
	for i := 0; i < len(d_file.cards); i++ {
		if (d_file.cards[i].value != d_game.cards[i].value) || (d_file.cards[i].suit != d_game.cards[i].suit) {
			t.Errorf("Card %d in Returned Deck is %s.", i, d_game.cards[i].value+" of "+d_game.cards[i].suit)
			t.Errorf("Card %d in File Deck is %s.", i, d_file.cards[i].value+" of "+d_file.cards[i].suit)
		}
	}
}
