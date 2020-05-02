package main

import "fmt"

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
	return false
}

func duplicate(state ShogiState) ShogiState {
	newBoard := make([][]string, len(state.board))
	for i := 0; i < len(newBoard); i++ {
		newBoard[i] = make([]string, len(state.board[i]))
		copy(newBoard[i], state.board[i])
	}
	newState := ShogiState{board: newBoard, parent: state.parent}
	return newState
}

func (state ShogiState) Succ() []ShogiState {
	var final []ShogiState
	for i := 0; i < len(state.pieces); i++ {
		//Scan through all pieces and appends all possible moves of all pieces to the final slice
	}

	return final
}

func getFirst(state ShogiState) Move {
	//This should loop to follow the parent pointer to get the first move that led to this chain
	return state
}

func (state ShogiState) Equal(s ShogiState) bool {
	if len(state.board) != len(s.board) {
		return false
	}
	for i := 0; i < len(state.board); i++ {
		if len(state.board[i]) != len(s.board[i]) {
			return false
		}
		for j := 0; j < len(state.board[i]); j++ {
			if state.board[i][j] != s.board[i][j] {
				return false
			}
		}
	}

	if len(state.pieces) != len(s.pieces) {
		return false
	}
	for i := 0; i < len(state.pieces); i++ {
		for j := 0; j < len(state.pieces); j++ {
			if (state.pieces[i].x == state.pieces[j].x) != (s.pieces[i].y == s.pieces[j].y) {
				return false
			}
		}
	}
	return true
}

func ExpectoMax(state ShogiState, depth int) (Move, error) {
	var final []ShogiState
	var fringe []ShogiState
	fringe = append(fringe, state)

	for {
		if len(fringe) == 0 {
			return Move{}, fmt.Errorf("Sad boi")
		}
		node := fringe[0]
		fringe = popList(fringe)

		if node.IsGoal() && len(fringe) == 0 {
			final = append(final, node)
			return getFirst(final[len(final)-1]), nil
		} else {
			kids := node.Succ()
			//Make some decisions here about the kids
			fringe = append(kids, fringe...)
			final = append(final, node)
		}

	}
}
