/******************/
/* Play Baccarat */
/*****************/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const winThreshold = 8
const finalRoundThreshold = 5

/*
getCardScore returns the score value associated with a given card.

The `card` argument is a string representation of a card, in the format "{value} of {suit}".
The function extracts the card value from the string by splitting it at the space character.
If the card value is numeric, it is parsed into an integer using the `strconv.Atoi` function
and stored in the `score` variable. If the card value is "Ace", the `score` is set to 1. If
the card value is not numeric and cannot be parsed into an integer, the `score` is set to 0.

The function then returns the `score` value, which represents the score associated with the
given card.
*/
func getCardScore(card string) int {
	card_value := strings.Split(card, " ")[0]
	score, err := strconv.Atoi(card_value)

	if card_value == "Ace" {
		score = 1
	} else if err != nil {
		score = 0
	}

	return score
}

/*
calculateScore calculates the score of a deck of cards and prints it to the console.

The method takes an additional `individual` argument, which is a string representing
the identifier or name associated with the deck. This argument is used to display the
score in the console output.

The method iterates through each card in the deck and accumulates the score by calling
the `getCardScore` function on each card. The scores are summed together. After all cards
are processed, the total score is modulo-divided by 10 to obtain the final score.

The method then prints the calculated score to the console, formatted as "{individual}
Score : {score}", where {individual} is the identifier passed as an argument, and {score}
is the calculated score.

Finally, the method returns the calculated score as an integer.
*/
func (d deck) calculateScore(individual string) int {
	score := 0
	for _, card := range d {
		score += getCardScore(card)
	}
	score = score % 10

	fmt.Printf("%s Score : %d\n\n", individual, score)
	return score
}

/*
getWinner determines the winner between a player and a banker based on their scores.

The function takes two integer arguments: `player_score` representing the score of the player,
and `banker_score` representing the score of the banker. The scores are compared to determine
the winner of the game.

If the `player_score` is greater than the `banker_score`, the function returns "Player" as the
winner. If the `banker_score` is greater than the `player_score`, the function returns "Banker"
as the winner. If the scores are equal, indicating a tie, the function returns "Tie".

The function does not handle any special cases or tie-breaking rules beyond a simple comparison
of the scores.

The function returns a string indicating the winner ("Player", "Banker", or "Tie").
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
*/
func play_baccarat() string {
	game_deck := newDeck()

	player_deck := deck{}
	banker_deck := deck{}

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

	player_deck.print("Player")
	player_score := player_deck.calculateScore("Player")

	banker_deck.print("Banker")
	banker_score := banker_deck.calculateScore("Banker")

	if (player_score >= winThreshold) || (banker_score >= winThreshold) {
		return getWinner(player_score, banker_score)
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

		// Determine if banker gets a third card
		banker_final = (banker_final) || (banker_score <= (finalRoundThreshold - 3))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 2)) && (getCardScore(player_deck[2]) != winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 1)) && (getCardScore(player_deck[2]) >= (banker_score - 2)) && (getCardScore(player_deck[2]) < winThreshold))
		banker_final = (banker_final) || ((banker_score == finalRoundThreshold) && (getCardScore(player_deck[2]) >= (banker_score - 1)) && (getCardScore(player_deck[2]) < winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold + 1)) && (getCardScore(player_deck[2]) >= banker_score) && (getCardScore(player_deck[2]) < winThreshold))
	}

	if banker_final {
		game_deck.deal(&banker_deck, 1)
	}

	/********************/
	/* Calculate Scores */
	/********************/

	player_deck.print("Player")
	player_score = player_deck.calculateScore("Player")

	banker_deck.print("Banker")
	banker_score = banker_deck.calculateScore("Banker")

	return getWinner(
		player_score,
		banker_score,
	)
}
