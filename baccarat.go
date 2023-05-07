/******************/
/* Play Baccarat */
/*****************/

package main

func play_baccarat() string {
	playing_deck := initDeck()

	player_deck := []string{}
	banker_deck := []string{}
	/***********/
	/* Round 1 */
	/***********/

	playing_deck, player_deck = addCards(playing_deck, player_deck, 1)
	playing_deck, banker_deck = addCards(playing_deck, banker_deck, 1)

	/***********/
	/* Round 2 */
	/***********/

	playing_deck, player_deck = addCards(playing_deck, player_deck, 1)
	playing_deck, banker_deck = addCards(playing_deck, banker_deck, 1)

	/********************/
	/* Calculate Scores */
	/********************/

	player_score := calculateScore(player_deck, "Player")
	banker_score := calculateScore(banker_deck, "Banker")
	if (player_score >= winThreshold) || (banker_score >= winThreshold) {
		if player_score > banker_score {
			return "Player"
		} else if banker_score > player_score {
			return "Opponent"
		} else {
			return "Tie"
		}
	}

	/***********/
	/* Round 3 */
	/***********/

	player_final := (player_score <= finalRoundThreshold)
	banker_final := false

	if !(player_final) {
		banker_final = (banker_score <= finalRoundThreshold)
	} else {
		playing_deck, player_deck = addCards(playing_deck, player_deck, 1)

		// Determine if Banker Gets a Third Card
		banker_final = (banker_final) || (banker_score <= (finalRoundThreshold - 3))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 2)) && (getCardScore(player_deck[2]) != winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold - 1)) && (getCardScore(player_deck[2]) >= (banker_score - 2)) && (getCardScore(player_deck[2]) < winThreshold))
		banker_final = (banker_final) || ((banker_score == finalRoundThreshold) && (getCardScore(player_deck[2]) >= (banker_score - 1)) && (getCardScore(player_deck[2]) < winThreshold))
		banker_final = (banker_final) || ((banker_score == (finalRoundThreshold + 1)) && (getCardScore(player_deck[2]) >= banker_score) && (getCardScore(player_deck[2]) < winThreshold))
	}

	if banker_final {
		playing_deck, banker_deck = addCards(playing_deck, banker_deck, 1)
	}

	/********************/
	/* Calculate Scores */
	/********************/

	player_score = calculateScore(player_deck, "Player")
	banker_score = calculateScore(banker_deck, "Banker")

	if player_score > banker_score {
		return "Player"
	} else if banker_score > player_score {
		return "Opponent"
	} else {
		return "Tie"
	}
}
