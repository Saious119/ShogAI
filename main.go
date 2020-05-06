package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
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

	currX--
	currY--
	finalX--
	finalY--

	for i := 0; i < len(state.pieces); i++ {
		if state.pieces[i].x == currX && state.pieces[i].y == currY {
			if newState.board[finalY][finalX] != "O" {
				for j := 0; j < len(newState.pieces); j++ {
					if newState.pieces[j].x == finalX && newState.pieces[j].y == finalY {
						newState.pieces[j].name = ""
						break
					}
				}
			}
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

func (state ShogiState) addPiece(coords []string) ShogiState {
	newState := duplicate(state)
	newX, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	newY, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}
	newPair := Pair{x: newX, y: newY, name: coords[2]}
	newState.pieces = append(newState.pieces, newPair)
	newState.board[newY][newX] = coords[2]
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
	searchDepth := 2
	var history []Move
	var oldData []byte
	var data []byte

	for {
		fmt.Println(state)
		m, err := MiniMax(state, player, searchDepth, history)
		if err != nil {
			panic(err)
		}
		fmt.Println(m)
		coords := strings.Split(m.String(), " ")
		ioutil.WriteFile("./NodeScriptShogAI/move.txt", []byte(m.String()), 0644)

		time.Sleep(1 * time.Second)

		for {
			data, err = ioutil.ReadFile("./NodeScriptShogAI/board.txt")
			if err != nil {
				fmt.Println("Failed to parse")
				continue
			}
			if string(data) == string(oldData) {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			if string(data) == "try again" {
				continue
			}
			fmt.Println(string(data))
			oldData = append([]byte{}, data...)
			break
		}
		state = state.updateCoords(coords)
		coords = strings.Split(string(data), " ")
		history = append(history, m)
		if len(coords) == 4 {
			state = state.updateCoords(coords)
		} else if len(coords) == 3 {
			state = state.addPiece(coords)
		}

		if state.IsGoal(player) {
			fmt.Println("We won!")
			break
		} else if state.IsGoal(((player + 1) % 2) + 1) {
			fmt.Println("We lost!")
		}
	}

}
