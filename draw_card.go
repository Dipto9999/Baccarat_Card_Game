/***************/
/* Draw a Card */
/***************/

package main

import (
	"math/rand"
	"time"
)

func getNumber(max_number int) int {
	// Seed the Random Number Generator.
	rand.Seed(time.Now().UnixNano())

	// Generate a Random Number Within Specified Range.
	return rand.Intn(max_number)
}

func generateCard(playing_deck []string) string {
	return playing_deck[getNumber(len(playing_deck))]
}
