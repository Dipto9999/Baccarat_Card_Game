# Baccarat Card Game

This is a command-line Baccarat card game written in Golang. In this game, you play against a computer dealer with the goal of having a hand value closer to 9 than the dealer's hand. The game uses a standard deck of cards where face cards (J, Q, K) and 10s have a value of zero, and Aces have a value of one.

## Table of Contents

- [Setup](#setup)
- [Running the Game](#running-the-game)
- [Playing the Game](#playing-the-game)
- [Code Structure](#code-structure)
- [Credit](#credit)

## Setup

To play the game, you need to have Golang installed on your system. To check if Golang is installed on your system, run the following command:

```
go version
```

If you don't have Golang installed, you can download and install it from the official Golang website: https://golang.org/dl/

## Running the Game

To run the game, navigate to the project directory and run the following command:

```
go run main.go baccarat.go deck.go
```

## Playing the Game

The game proceeds as follows:

1. Two cards are dealt to both the player and the dealer.
2. The player's hand value is displayed.
3. The player can choose to draw another card or stand.
4. If the player chooses to draw, another card is dealt, and the hand value is updated.
5. If the player's hand value exceeds 9, the hand value is reset to the last digit of the sum.
6. If the player chooses to stand, the dealer's hand is revealed, and the winner is determined based on who has the hand closest to 9.
7. If the dealer's hand value is less than 6, the dealer draws another card.
8. If the dealer's hand value exceeds 9, the hand value is reset to the last digit of the sum.
9. If the player's hand value is equal to the dealer's hand value, the game ends in a tie.

## Code Structure

The code is structured into the following files:

- `main.go`: contains the main function.
- `baccarat.go`: contains the main game logic and rules.
- `deck.go`: contains the deck creation and shuffling logic.

## Credit

This Baccarat card game was created with inspiration from the following sources:

- A Digital Systems Design course (CPEN 311) at the University of British Columbia, which implements the game on an FPGA with Verilog.
- The Udemy Course on Go: The Complete Developer's Guide (Golang) by Stephen Grider.

This README was generated by ChatGPT.