package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

type Move struct {
	curr  Pair
	final Pair
}

type ShogiState struct {
	board  [][]string
	pieces []Pair
	parent *ShogiState
}

func popList(list []ShogiState) []ShogiState {
	return list[1:]
}

func (state ShogiState) IsGoal() bool {
	//please work on
	return false
}

func duplicate(state ShogiState) ShogiState {
	newBoard := make([][]string, len(state.board))
	for i := 0; i < len(newBoard); i++ {
		newBoard[i] = make([]string, len(state.board[i]))
		copy(newBoard[i], state.board[i])
	}
	newPieces := make([]Pair, len(state.pieces))
	copy(newPieces, state.pieces)
	newState := ShogiState{board: newBoard, pieces: newPieces, parent: state.parent}
	return newState
}

func (state ShogiState, player int) Succ() []ShogiState {
	var final []ShogiState
	for i := 0; i < len(state.pieces); i++ {
		//Scan through all pieces and appends all possible moves of all pieces to the final slice
	}

	return final
}

func diff(s1, s2 ShogiState) (Move, error) {
	var m Move
	for i := 0; i < len(s1.pieces); i++ {
		if s1.pieces[i].x != s2.pieces[i].x || s1.pieces[i].y != s2.pieces[i].y {
			m.curr = s1.pieces[i]
			m.final = s2.pieces[i]
			return m, nil
		}
	}
	return m, fmt.Errorf("No diff! You're a liar! You promiced me!")
}

func getFirstMove(state ShogiState) Move {
	//Totally untested and highly dangerous
	curr := state
	final := curr
	for curr.parent != nil {
		final = curr
		curr = *(curr.parent)
	}
	m, err := diff(curr, final)
	if err != nil {
		panic("No move made, even though moves made! Someone call the Navy!")
	}

	return m
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func auxMiniMax(state ShogiState, player, depth int, max bool) int {
	if state.IsGoal() {
		//Oh Mowie Wowie!
		if max {
			return 9999
		}
		return -9999
	}

	if depth == 0 {
		if player == 1 {
			return h1(state.board)
		} else {
			return h2(state.board)
		}
	} else if max {
		val := 0
		kids := state.Succ()
		for i := 0; i < len(kids); i++ {
			val = Max(val, auxMiniMax(kids[i], player, depth-1, false))
		}
		return val
	} else {
		val := 0
		kids := state.Succ()
		for i := 0; i < len(kids); i++ {
			val = Min(val, auxMiniMax(kids[i], player, depth-1, true))
		}
		return val
	}
}

func MiniMax(state ShogiState, player, depth int) (Move, error) {
	//Totally untested, und highly dangerous! (waiting for Succ)
	if state.IsGoal() {
		return Move{}, fmt.Errorf("You won dufus! Email all your friends!")
	}

	kids := state.Succ()
	var vals []int
	for i := 0; i < len(kids); i++ {
		val := auxMiniMax(kids[i], player, depth-1, false)
		vals = append(vals, val)
	}

	maximum := 0
	for i := 0; i < len(vals); i++ {
		if vals[i] > maximum {
			maximum = i
		}
	}

	m, err := diff(state, kids[maximum])
	if err != nil {
		panic("No move is max, this shouldn't be possible because maximum := 0")
	}
	return m, nil
}

func OwnsPiece(piece string, playerNum int) bool {
	strPlayerNum := strconv.Itoa(playerNum)
	if strings.Contains(piece, strPlayerNum) {
		return true
	}
	return false
}
