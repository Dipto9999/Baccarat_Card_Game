/******************/
/* Play Baccarat */
/*****************/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

func readFromFile(filename string) deck {
	b_slice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(b_slice), "\n"))

}
