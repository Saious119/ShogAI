package main

import "fmt"

func InitBoard() [][]string { //creates a new board for a new game and prints it
	NewBoard := [][]string{
		{"L1", "N1", "S1", "G1", "K1", "G1", "S1", "N1", "L1"},
		{"O", "R1", "O", "O", "O", "O", "O", "B1", "O"},
		{"P1", "P1", "P1", "P1", "P1", "P1", "P1", "P1", "P1"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "0"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "0"},
		{"O", "O", "O", "O", "O", "O", "O", "O", "0"},
		{"P2", "P2", "P2", "P2", "P2", "P2", "P2", "P2", "P2"},
		{"O", "B2", "O", "O", "O", "O", "O", "R2", "O"},
		{"L2", "N2", "S2", "G2", "K2", "G2", "S2", "N2", "L2"},
	}
	//printBoard(NewBoard)
	return NewBoard
}

func PrintBoard(board [][]string) { //prints the given board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(" ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func InitPieces() []Pair {
	var pieces []Pair
	pair := Pair{0, 0, "L1"}
	pieces = append(pieces, pair)
	pair = Pair{8, 0, "L1"}
	pieces = append(pieces, pair)
	pair = Pair{0, 8, "L2"}
	pieces = append(pieces, pair)
	pair = Pair{8, 8, "L2"}
	pieces = append(pieces, pair)

	pair = Pair{1, 0, "N1"}
	pieces = append(pieces, pair)
	pair = Pair{7, 0, "N1"}
	pieces = append(pieces, pair)
	pair = Pair{1, 8, "N2"}
	pieces = append(pieces, pair)
	pair = Pair{7, 8, "N2"}
	pieces = append(pieces, pair)

	pair = Pair{2, 0, "S1"}
	pieces = append(pieces, pair)
	pair = Pair{6, 0, "S1"}
	pieces = append(pieces, pair)
	pair = Pair{2, 8, "S2"}
	pieces = append(pieces, pair)
	pair = Pair{6, 8, "S2"}
	pieces = append(pieces, pair)

	pair = Pair{3, 0, "G1"}
	pieces = append(pieces, pair)
	pair = Pair{5, 0, "G1"}
	pieces = append(pieces, pair)
	pair = Pair{3, 8, "G2"}
	pieces = append(pieces, pair)
	pair = Pair{5, 8, "G2"}
	pieces = append(pieces, pair)

	pair = Pair{4, 0, "K1"}
	pieces = append(pieces, pair)
	pair = Pair{4, 8, "K2"}
	pieces = append(pieces, pair)

	pair = Pair{1, 1, "R1"}
	pieces = append(pieces, pair)
	pair = Pair{1, 7, "R2"}
	pieces = append(pieces, pair)

	pair = Pair{7, 1, "B1"}
	pieces = append(pieces, pair)
	pair = Pair{7, 7, "B2"}
	pieces = append(pieces, pair)

	for i := 0; i < 9; i++ {
		pair = Pair{i, 2, "P1"}
		pieces = append(pieces, pair)
	}

	for i := 0; i < 9; i++ {
		pair = Pair{i, 6, "P2"}
		pieces = append(pieces, pair)
	}

	return pieces
}
