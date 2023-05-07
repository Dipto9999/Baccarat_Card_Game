/***********************/
/* Score Baccarat Hand */
/***********************/

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

func calculateScore(effective_deck []string, individual string) int {
	for _, current_card := range effective_deck {
		fmt.Printf("%s : %s\n", individual, current_card)
	}

	effective_score := 0
	for _, current_card := range effective_deck {
		effective_score += getCardScore(current_card)
	}
	effective_score = effective_score % 10

	fmt.Printf("%s Score : %d\n\n", individual, effective_score)
	return effective_score
}
