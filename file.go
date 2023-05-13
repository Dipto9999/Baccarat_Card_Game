/*************/
/* File I/O */
/************/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
saveToFile saves the deck of cards to a file.

The method takes a `filename` argument, which is a string representing the name of the file
to which the deck should be saved.

The method converts the deck into a byte slice using the `toByteSlice` method and writes the
byte slice to the specified file using the `ioutil.WriteFile` function. The file is created
if it does not exist, and if it already exists, its contents are overwritten.

The method returns an error if there is an issue writing the file, or `nil` if the file is
successfully written.
*/
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

/*
readFromFile reads a deck of cards from a file. The file must contain a single card per line,
with each line terminated by a newline character ("\n").
*/
func readFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile("game_deck.txt")
	if err != nil {
		fmt.Sprintf("Error Reading File: %s\n", err)
	}

	return deck(deck(strings.Split(string(byteSlice), "\n")))
}
