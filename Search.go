package main

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

func (state ShogiState) atGoal() bool {
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

	}

	return final
}

func reverseChain(state ShogiState) []ShogiState {
	var final []ShogiState
	curr := state
	for curr.parent != nil {
		final = append(final, curr)
		curr = *(curr.parent)
	}
	if curr.parent == nil {
		final = append(final, curr)
	}

	var ordered []ShogiState
	for i := len(final) - 1; i >= 0; i-- {
		ordered = append(ordered, final[i])
	}
	return ordered
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

func ExpectoMax(state ShogiState, depth int) Move {
	// var final []ShogiState
	var fringe []ShogiState
	fringe = append(fringe, state)

	return Move{}
}
