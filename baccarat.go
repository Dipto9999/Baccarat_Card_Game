/******************/
/* Play Baccarat */
/*****************/

package main

import (
	"fmt"
	"strconv"
)

const winThreshold = 8
const finalRoundThreshold = 5

/*
getScore returns the score of the card.

If the card's value is "Ace", the score is 1.
If the card's value can be converted to an integer, the score is that integer.
Otherwise, the score is 0.

Receiver:
- c: card object to get the score for.

Returns:
- card score as an integer.
*/
func (c *card) getScore() int {
	score, err := strconv.Atoi(c.value)

	if c.value == "Ace" {
		score = 1
	} else if err != nil {
		score = 0
	}

	return score
}

/*
calculateScore calculates the score of the deck and assigns it to the deck's score field.

The score is calculated by summing up the scores of all cards in the deck,
then taking the remainder when divided by 10.

Prints the owner of the deck and the calculated score to the console.

Receiver:
- d: pointer to the deck to calculate the score for.

Returns:
- score of the deck as an integer.
*/
func (d *deck) calculateScore() int {
	score := 0
	for _, card := range d.cards {
		score += card.getScore()
	}
	d.score = score % 10

	fmt.Printf("%s Score : %d\n", d.owner, d.score)

	return d.score
}

/*
foundWinner determines whether a winner has been found.
If the game is in the second round, it is based on whether
one of the scores is greater than or equal to the winThreshold.
If the game is in the third round, the game is over.

Parameters:
- d_1: pointer to the first deck to compare.
- d_2: pointer to the second deck to compare.

Returns:
- boolean value indicating whether a winner has been found.
*/
func foundWinner(d_1 *deck, d_2 *deck) bool {

	curr_round := func(a int, b int) int {
		// Return the greater of the two values.
		if a >= b {
			return a
		} else {
			return b
		}
	}(len((*d_1).cards), len((*d_2).cards))

	// Must play at least 2 rounds.
	if curr_round < 2 {
		return false
	}

	fmt.Printf("\nAfter Round %d : \n\n", curr_round)
	(*d_1).print()
	(*d_2).print()

	(*d_1).calculateScore()
	(*d_2).calculateScore()

	if curr_round == 2 {
		return ((*d_1).score >= winThreshold) || ((*d_2).score >= winThreshold)
	} else {
		return true
	}
}

/*
getWinner determines the winner based on the scores of the two decks.

If there is a clear winner, the name of the winner is returned.
Otherwise, "Tie" is returned.

Parameters:
- d_1: the first deck to compare.
- d_2: the second deck to compare.

Returns
- name of the winner as a string.
*/
func getWinner(d_1 deck, d_2 deck) string {
	if d_1.score > d_2.score {
		return d_1.owner
	} else if d_2.score > d_1.score {
		return d_2.owner
	} else {
		return "Tie"
	}
}

/*
play_baccarat simulates a game of Baccarat and determines the winner.

The function initializes a new deck using the `newDeck` function and creates empty decks for
the player and the banker. The game proceeds in multiple rounds, with each round involving
dealing cards to the player and the banker from the game deck.

Round 1:
- One card is dealt to the player.
- One card is dealt to the banker.

Round 2:
- One card is dealt to the player.
- One card is dealt to the banker.

The scores of the player and the banker are calculated using the `calculateScore` method of
their respective decks. The scores are printed to the console along with the identifiers
("Player" or "Banker").

If either the player's score or the banker's score reaches or exceeds the `winThreshold`, the
game ends and the winner is determined using the `getWinner` function. The winner's name
("Player", "Banker", or "Tie") is returned as a string.

Round 3:
- If the player's score is less than or equal to the `finalRoundThreshold`, the player receives
a third card.
  - The banker will draw a third card if the banker's score is `winThreshold` - 2 and
    the player's third card value is greater than or equal to the banker's score and less than the `winThreshold`.
  - The banker will draw a third card if the banker's score is `winThreshold` - 3 and
    the player's third card value is greater than or equal to 4 and less than `winThreshold`.
  - The banker will draw a third card if the banker's score is `winThreshold` - 4 and
    the player's third card value is greater than or equal to 2 and less than `winThreshold`.
  - The banker will draw a third card if the banker's score is `winThreshold` - 5 and
    the player's third card value is not equal to `winThreshold`.
  - The banker will draw a third card if the banker's score is less than or equal to 2.

- The banker will draw a third card in spite of the above if the banker's score is less than or\
equal to `finalRoundThreshold`.

The scores of the player and the banker are calculated again, and the results are printed to the
console. The winner is determined using the `getWinner` function, considering the updated scores.
The winner's name ("Player", "Banker", or "Tie") is returned as a string.

Returns:
- name of the winner as a string.
*/
func play_baccarat() string {
	game_deck := newDeck(5, "Game Deck")
	player_deck := deck{owner: "Player"}
	banker_deck := deck{owner: "Banker"}

	/***********/
	/* Round 1 */
	/***********/

	game_deck.deal(&player_deck, 1)
	game_deck.deal(&banker_deck, 1)

	/***********/
	/* Round 2 */
	/***********/

	game_deck.deal(&player_deck, 1)
	game_deck.deal(&banker_deck, 1)

	/********************/
	/* Calculate Scores */
	/********************/

	if foundWinner(&player_deck, &banker_deck) {
		return getWinner(player_deck, banker_deck)
	}

	/***********/
	/* Round 3 */
	/***********/

	// Determine if player and banker get a third card.
	var banker_final bool
	if player_deck.score <= finalRoundThreshold {
		game_deck.deal(&player_deck, 1)
		third_score := player_deck.cards[2].getScore()

		switch banker_deck.score {
		case (winThreshold - 1):
			banker_final = false
		case (winThreshold - 2):
			banker_final = (third_score >= banker_deck.score) &&
				(third_score < winThreshold)
		case (winThreshold - 3):
			banker_final = (third_score >= banker_deck.score-1) &&
				(third_score < winThreshold)
		case (winThreshold - 4):
			banker_final = (third_score >= banker_deck.score-2) &&
				(third_score < winThreshold)
		case (winThreshold - 5):
			banker_final = (third_score != winThreshold)
		case 2, 1, 0:
			banker_final = true
		default:
			banker_final = false
		}
	}

	banker_final = banker_final || (banker_deck.score <= finalRoundThreshold)
	if banker_final {
		game_deck.deal(&banker_deck, 1)
	}

	/********************/
	/* Calculate Scores */
	/********************/

	foundWinner(&player_deck, &banker_deck)
	return getWinner(player_deck, banker_deck)
}
