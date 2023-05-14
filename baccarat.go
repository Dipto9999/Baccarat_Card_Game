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
calculateScore calculates the score of the deck.

The score is calculated by summing up the scores of all cards in the deck,
then taking the remainder when divided by 10.

Prints the owner of the deck and the calculated score to the console.

Receiver:
- d: deck object to calculate the score for.

Returns:
- calculated score of the deck as an integer.
*/
func (d deck) calculateScore() int {
	score := 0
	for _, card := range d.cards {
		score += card.getScore()
	}
	score = score % 10

	fmt.Printf("%s Score : %d\n\n", d.owner, score)
	return score
}

/*
getWinner determines the winner based on the player and banker scores.

If the player's score is greater than the banker's score, "Player" is returned.
If the banker's score is greater than the player's score, "Banker" is returned.
Otherwise, "Tie" is returned.

Parameters:
- player_score: score of the player as an integer.
- banker_score: score of the banker as an integer.

Returns:
- name of the winner as a string.
*/
func getWinner(player_score int, banker_score int) string {
	if player_score > banker_score {
		return "Player"
	} else if banker_score > player_score {
		return "Banker"
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
    an additional card.
  - If the player's score is greater than the `finalRoundThreshold`, the banker determines whether
    to draw a third card based on specific rules and the player's third card, if any.
  - The banker may draw a third card if the score is less than or equal to `finalRoundThreshold - 3`.
  - The banker may draw a third card if the score is `finalRoundThreshold - 2` and the player's
    third card's score is not equal to the `winThreshold`.
  - The banker may draw a third card if the score is `finalRoundThreshold - 1` and the player's
    third card's score is greater than or equal to `banker_score - 2` and less than `winThreshold`.
  - The banker may draw a third card if the score is `finalRoundThreshold` and the player's
    third card's score is greater than or equal to `banker_score - 1` and less than `winThreshold`.
  - The banker may draw a third card if the score is `finalRoundThreshold + 1` and the player's
    third card's score is greater than or equal to `banker_score` and less than `winThreshold`.
  - Otherwise, the banker does not draw a third card.

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

	player_deck.print()
	player_score := player_deck.calculateScore()

	banker_deck.print()
	banker_score := banker_deck.calculateScore()

	if (player_score >= winThreshold) || (banker_score >= winThreshold) {
		return getWinner(
			player_score,
			banker_score,
		)
	}

	/***********/
	/* Round 3 */
	/***********/

	player_final := (player_score <= finalRoundThreshold)
	banker_final := false

	if !(player_final) {
		banker_final = (banker_score <= finalRoundThreshold)
	} else {
		game_deck.deal(&player_deck, 1)

		// Determine if banker gets a third card.
		banker_final = (banker_final) || (banker_score <= (finalRoundThreshold - 3))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 2)) && (player_deck.cards[2].getScore() != winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 1)) && (player_deck.cards[2].getScore() >= (banker_score - 2)) && (player_deck.cards[2].getScore() < winThreshold))
		banker_final = (banker_final) || ((banker_score == finalRoundThreshold) && (player_deck.cards[2].getScore() >= (banker_score - 1)) && (player_deck.cards[2].getScore() < winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold + 1)) && (player_deck.cards[2].getScore() >= banker_score) && (player_deck.cards[2].getScore() < winThreshold))
	}

	if banker_final {
		game_deck.deal(&banker_deck, 1)
	}

	/********************/
	/* Calculate Scores */
	/********************/

	player_deck.print()
	player_score = player_deck.calculateScore()

	banker_deck.print()
	banker_score = banker_deck.calculateScore()

	return getWinner(
		player_score,
		banker_score,
	)
}
