# Baccarat Card Game

This is a command-line Baccarat card game written in **Golang**. The game involves playing against a computer banker with the objective of obtaining a hand value closer to *9* than the banker's hand. A standard deck of cards is used, where numeric cards hold their face value, face cards (*J*, *Q*, *K*) are valued at 0, and *Aces* are valued at 1.

## Table of Contents

- [Setup](#setup)
- [Running the Game](#running-the-game)
- [Playing the Game](#playing-the-game)
- [Code Structure](#code-structure)
- [Tests](#tests)
- [Credit](#credit)

## Setup

To play the game, ensure that **Golang** is installed on your system. You can check if **Golang** is installed by running the following command:

```
go version
```

If **Golang** is not installed, you can download and install it from the official **Golang** website: [https://golang.org/dl/](https://golang.org/dl/)

## Running the Game

To run the game, navigate to the project directory and execute the following command:

```
go run main.go baccarat.go deck.go file.go
```

## Playing the Game

The game follows these steps:

1. Two cards are dealt to both the player and the dealer.
2. The scores are determined by the ones digit of the hand values.
3. The player can draw an additional card if their score is less than or equal to *8*.
4. Based on the current scores, the dealer may draw a third card.
5. If the hand value of the player or dealer exceeds *9*, the value is reset to the last digit of the sum.
6. If the player's hand value is equal to the dealer's hand value, the game results in a tie.

## Code Structure

The project is structured into the following files:

- `main.go`: contains the main function.
- `baccarat.go`: includes the main game logic and rules.
- `deck.go`: handles the deck creation and shuffling logic.
- `file.go`: manages the file reading and writing operations.
- `deck_test.go`: contains test functions to verify the correctness of the deck logic.

## Tests

The `deck_test.go` file includes test cases to ensure the accuracy of the deck logic. Here's an overview of the test cases:

- `TestInitDeck`: verifies the correct initialization of the deck with 56 cards, and checks the first and last cards.
- `TestNewDeck`: ensures that the returned deck and the deck read from the file both consist of 56 cards and match each other.
- `TestShuffle`: validates the proper shuffling of the deck by comparing it with the unshuffled deck.
- `TestDeal`: checks if the specified number of cards are dealt correctly and if the decks are updated accordingly.

To run the test suite, use the following command:

```
go test -v
```

This command executes all the test cases in the `deck_test.go` file and provides detailed information about the test results.

## Credit

This Baccarat card game was created with inspiration from the following sources:

- The **CPEN 311 - Digital Systems Design** course in the **University of British Columbia Electrical and Computer Engineering** Undergraduate program, which requires building a Baccarat engine in **Verilog**.
- The **Udemy** Course on **Go: The Complete Developer's Guide (Golang)** by Stephen Grider.

*This README was generated with the help of OpenAI's ChatGPT.*