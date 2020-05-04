package main

func Pawn(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) { //check if its a valid move
		state.board[NewX][NewY] = "P1" //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "P1+"
			state.pieces[i].name = "P1+"
		}
	}
	return state
}

func Lance(state ShogiState, player int, NewX int, NewY int, i int) ShogiState {
	if IsValid(state.board, NewX, NewY, player) { //check if its a valid move
		state.board[NewX][NewY] = "L1" //update board
		state.board[state.pieces[i].x][state.pieces[i].y] = "O"
		state.pieces[i].x = NewX //update piece
		state.pieces[i].y = NewY
		if CheckPromotion(NewY, state.pieces[i].name) {
			state.board[NewX][NewY] = "L1+"
			state.pieces[i].name = "L1+"
		}
	}
	return state
}
