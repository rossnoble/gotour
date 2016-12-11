package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// Loop over rows of board and print
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// => _ _ _
	// => _ _ _
	// => _ _ _

	// Players take their turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// Print the outcome of the game
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// => X _ X
	// => O _ X
	// => _ _ O
}
