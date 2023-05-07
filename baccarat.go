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

func (cards deck) calculateScore(individual string) int {
	score := 0
	for _, card := range cards {
		score += getCardScore(card)
	}
	score = score % 10

	fmt.Printf("%s Score : %d\n", individual, score)
	return score
}

func getWinner(player_score int, banker_score int) string {
	if player_score > banker_score {
		return "Player"
	} else if banker_score > player_score {
		return "Banker"
	} else {
		return "Tie"
	}
}

func play_baccarat() string {
	game_deck := newDeck()

	player_deck := deck{}
	banker_deck := deck{}

	/***********/
	/* Round 1 */
	/***********/

	player_deck.addCards(&game_deck, 1)
	banker_deck.addCards(&game_deck, 1)

	/***********/
	/* Round 2 */
	/***********/

	player_deck.addCards(&game_deck, 1)
	banker_deck.addCards(&game_deck, 1)

	/********************/
	/* Calculate Scores */
	/********************/

	player_deck.print("Player")
	player_score := player_deck.calculateScore("Player")

	banker_deck.print("Banker")
	banker_score := banker_deck.calculateScore("Banker")

	if (player_score >= winThreshold) || (banker_score >= winThreshold) {
		return getWinner(player_score, banker_score)
	} else {
		fmt.Printf("\n")
	}

	/***********/
	/* Round 3 */
	/***********/

	player_final := (player_score <= finalRoundThreshold)
	banker_final := false

	if !(player_final) {
		banker_final = (banker_score <= finalRoundThreshold)
	} else {
		player_deck.addCards(&game_deck, 1)

		// Determine if Banker Gets a Third Card
		banker_final = (banker_final) || (banker_score <= (finalRoundThreshold - 3))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 2)) && (getCardScore(player_deck[2]) != winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 1)) && (getCardScore(player_deck[2]) >= (banker_score - 2)) && (getCardScore(player_deck[2]) < winThreshold))
		banker_final = (banker_final) || ((banker_score == finalRoundThreshold) && (getCardScore(player_deck[2]) >= (banker_score - 1)) && (getCardScore(player_deck[2]) < winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold + 1)) && (getCardScore(player_deck[2]) >= banker_score) && (getCardScore(player_deck[2]) < winThreshold))
	}

	if banker_final {
		banker_deck.addCards(&game_deck, 1)
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
