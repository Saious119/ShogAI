package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	board := InitBoard()
	// PrintBoard(board)
	state := ShogiState{
		board:  board,
		pieces: InitPieces(),
		parent: nil,
	}
	// fmt.Println(state)

	// bFile, err := os.Open("./NodeScriptShogAI/board.txt")
	for {
		fmt.Println(state)
		m, err := MiniMax(state, 2, 1)
		if err != nil {
			panic(err)
		}
		fmt.Println(m)
		ioutil.WriteFile("./NodeScriptShogAI/move.txt", []byte(m.String()), 0644)
		time.Sleep(30 * time.Second)
	}

}
