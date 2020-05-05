package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func (state ShogiState) updateCoords(coords []string) ShogiState {
	newState := duplicate(state)
	currX, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	currY, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}
	finalX, err := strconv.Atoi(coords[2])
	if err != nil {
		panic(err)
	}
	finalY, err := strconv.Atoi(coords[3])
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(state.pieces); i++ {
		if state.pieces[i].x == currX && state.pieces[i].y == currY {
			piece := newState.pieces[i].name
			newState.board[newState.pieces[i].y][newState.pieces[i].x] = "O"
			newState.board[finalY][finalX] = piece //update board
			newState.pieces[i].x = finalX          //update piece
			newState.pieces[i].y = finalY
			if !strings.Contains(piece, "+") {
				if !strings.Contains(piece, "K") {
					if CheckPromotion(finalY, newState.pieces[i].name) {
						newState.board[finalY][finalX] = piece + "+"
						newState.pieces[i].name = piece + "+"
					}
				}
			}
			break
		}
	}

	return newState
}

func main() {
	board := InitBoard()
	state := ShogiState{
		board:  board,
		pieces: InitPieces(),
		parent: nil,
	}
	player := 2
	searchDepth := 1

	for {
		fmt.Println(state)
		m, err := MiniMax(state, player, searchDepth)
		if err != nil {
			panic(err)
		}
		fmt.Println(m)
		ioutil.WriteFile("./NodeScriptShogAI/move.txt", []byte(m.String()), 0644)

		data, err := ioutil.ReadFile("./NodeScriptShogAI/board.txt")
		if err != nil {
			panic(err)
		}
		coords := strings.Split(string(data), " ")
		state = state.updateCoords(coords)
		if state.IsGoal(player) {
			fmt.Println("We won!")
			break
		}
	}

}
