package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	x    int
	y    int
	name string
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

func Succ(state ShogiState, player int) []ShogiState {
	var final []ShogiState
	for i := 0; i < len(state.pieces); i++ {
		//Scan through all pieces and appends all possible moves of all pieces to the final slice
		if OwnsPiece(state.pieces[i].name, player) {
			switch state.pieces[i].name {
			case "P1", "P2":
				NewX := state.pieces[i].x + 0
				var NewY int
				if strings.Contains(state.pieces[i].name, "1") {
					NewY = state.pieces[i].y - 1
				}
				if strings.Contains(state.pieces[i].name, "2") {
					NewY = state.pieces[i].y + 1
				}
				NewState := Pawn(state, player, NewX, NewY, i) //gives either a new state or if its invalid the same state
				final = append(final, NewState)
			case "L1":
				for j := state.pieces[i].y; j < len(state.board[0]); j++ {
					NewX := state.pieces[i].x
					NewY := j
					NewState := Lance(state, player, NewX, NewY, i)
					final = append(final, NewState)
					if state.board[NewX][NewY] != "O" {
						break
					}
				}
			case "L2":
				for j := state.pieces[i].y; j >= 0; j-- {
					NewX := state.pieces[i].x
					NewY := j
					NewState := Lance(state, player, NewX, NewY, i)
					final = append(final, NewState)
					if state.board[NewX][NewY] != "O" {
						break
					}
				}
			case "N1", "N2":
				NewX := state.pieces[i].x - 1
				var NewY int
				if strings.Contains(state.pieces[i].name, "1") {
					NewY = state.pieces[i].y - 2
				}
				if strings.Contains(state.pieces[i].name, "2") {
					NewY = state.pieces[i].y + 2
				}
				NewState := Knight(state, player, NewX, NewY, i)
				final = append(final, NewState)
				NewY = state.pieces[i].x + 1
				NewState = Knight(state, player, NewX, NewY, i)
				final = append(final, NewState)
			case "S1", "S2":
				var moves []int
				if strings.Contains(state.pieces[i].name, "1") {
					moves = []int{-1, -1, 0, -1, 1, -1, -1, 1, 1, 1}
				}
				if strings.Contains(state.pieces[i].name, "1") {
					moves = []int{-1, 1, 0, 1, 1, 1, -1, -1, 1, -1}
				}
				for j := 0; j <= len(moves)-2; j += 2 {
					NewX := state.pieces[j].x + moves[j]
					NewY := state.pieces[j].y + moves[j+1]
					NewState := Silver(state, player, NewX, NewY, i)
					final = append(final, NewState)
				}
			case "G1", "G2":
				var moves []int
				if strings.Contains(state.pieces[i].name, "1") {
					moves = []int{-1, -1, 0, -1, 1, -1, 1, 0, -1, 0, 0, 1}
				}
				if strings.Contains(state.pieces[i].name, "1") {
					moves = []int{-1, 1, 0, 1, 1, 1, 0, -1, 1, 0, 0, -1}
				}
				for j := 0; j <= len(moves)-2; j += 2 {
					NewX := state.pieces[j].x + moves[j]
					NewY := state.pieces[j].y + moves[j+1]
					NewState := Gold(state, player, NewX, NewY, i)
					final = append(final, NewState)
				}
			case "B1", "B2":
				for j := 0; j < len(state.board); j++ {
					NewX := state.pieces[i].x + j
					NewY := state.pieces[i].y + j
					NewState := Bishop(state, player, NewX, NewY, i)
					final = append(final, NewState)
					if state.board[NewX][NewY] != "O" {
						break
					}
				}
				for j := len(state.board); j >= 0; j-- {
					NewX := state.pieces[i].x - j
					NewY := state.pieces[i].y - j
					NewState := Bishop(state, player, NewX, NewY, i)
					final = append(final, NewState)
					if state.board[NewX][NewY] != "O" {
						break
					}
				}
			case "R1", "R2":
				for j := 0; j < len(state.board); j++ {
					NewX := j
					NewY := state.pieces[i].y
					NewState := Rook(state, player, NewX, NewY, i)
					final = append(final, NewState)
					if state.board[NewX][NewY] != "O" {
						break
					}
				}
				for g := 0; g < len(state.board[0]); g++ {
					NewX := state.pieces[i].x
					NewY := g
					NewState := Rook(state, player, NewX, NewY, i)
					final = append(final, NewState)
					if state.board[NewX][NewY] != "O" {
						break
					}
				}
			case "K1", "K2":
				var moves []int
				moves = []int{-1, -1, 0, -1, 1, -1, -1, 1, 1, 1, 1, 0, -1, 0, 0, 1}
				for j := 0; j <= len(moves)-2; j += 2 {
					NewX := state.pieces[j].x + moves[j]
					NewY := state.pieces[j].y + moves[j+1]
					NewState := King(state, player, NewX, NewY, i)
					final = append(final, NewState)
				}
			}
		}
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
		kids := Succ(state, player)
		for i := 0; i < len(kids); i++ {
			val = Max(val, auxMiniMax(kids[i], player, depth-1, false))
		}
		return val
	} else {
		val := 0
		kids := Succ(state, player)
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

	kids := Succ(state, player)
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

func IsValid(board [][]string, NewX int, NewY int, player int) bool {
	if strings.Contains(board[NewX][NewY], "K") {
		return false
	}
	StrPlayer := strconv.Itoa(player)
	if strings.Contains(board[NewX][NewY], StrPlayer) {
		return false
	}
	if NewX > len(board[0]) || NewX < 0 {
		return false
	}
	if NewY > len(board) || NewY < 0 {
		return false
	}
	return true
}

func CheckPromotion(Newy int, piece string) bool {
	if Newy > 5 {
		switch piece {
		case "P1":
			return true
		case "L1":
			return true
		case "N1":
			return true
		case "S1":
			return true
		case "B1":
			return true
		case "R1":
			return true
		}
	}
	if Newy < 4 {
		switch piece {
		case "P2":
			return true
		case "L2":
			return true
		case "N2":
			return true
		case "S2":
			return true
		case "B2":
			return true
		case "R2":
			return true
		}
	}
	return false
}
