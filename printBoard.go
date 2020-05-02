package main

import "fmt"

func PrintBoard(board [][]string) { //prints the given board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(" ", board[i][j])
		}
		fmt.Println()
	}
}
