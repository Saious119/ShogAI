package main

import "strconv"

//validates movements and make changes, then sends back a changed state
func Pawn(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) { //check if its a valid move
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "P" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "P" + StrPlayer + "+"
			state.pieces[i].name = "P" + StrPlayer + "+"
		}
	}
	return state
}

func Lance(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) { //check if its a valid move
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "L" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "L" + StrPlayer + "+"
			state.pieces[i].name = "L" + StrPlayer + "+"
		}
	}
	return state
}

func Knight(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) {
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "N" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "N" + StrPlayer + "+"
			state.pieces[i].name = "N" + StrPlayer + "+"
		}
	}
	return state
}

func Silver(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) {
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "S" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "S" + StrPlayer + "+"
			state.pieces[i].name = "S" + StrPlayer + "+"
		}
	}
	return state
}

func Gold(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) {
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "G" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
	}
	return state
}

func Bishop(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) {
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "B" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "B" + StrPlayer + "+"
			state.pieces[i].name = "B" + StrPlayer + "+"
		}
	}
	return state
}

func Rook(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) {
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "B" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "B" + StrPlayer + "+"
			state.pieces[i].name = "B" + StrPlayer + "+"
		}
	}
	return state
}

func King(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) {
		StrPlayer := strconv.Itoa(player)
		state.board[NewX][NewY] = "K" + StrPlayer //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
	}
	return state
}
