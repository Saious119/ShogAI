package main

import "fmt"

func main() {
	board := InitBoard()
	PrintBoard(board)
	state := ShogiState{
		board:  board,
		pieces: InitPieces(),
		parent: nil,
	}

	m, err := MiniMax(state, 1, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)

}
