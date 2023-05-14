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
saveToFile converts a deck of cards to a byte slice and writes it to a file.

Receiver:
- d: deck to convert to a byte slice and write to a file.

Parameters:
- filename: name of the file to write to.

Returns:
- error if any.
*/
func (d *deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

/*
readFromFile reads contents from a file and converts it to a deck of cards.

Parameters:
- filename: name of the file to read from.

Returns:
- deck struct representing the contents of the file.
*/
func readFromFile(filename string) deck {
	byte_slice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error Reading File: %s\n", err)
	}

	string_slice := strings.Split(string(byte_slice), "\n")

	d := deck{}
	for i := 0; i < len(string_slice); i++ {
		c := strings.Split(string_slice[i], " of ")
		d.cards = append(
			d.cards, card{
				value: c[0],
				suit:  c[1],
			},
		)
	}

	return d
}
