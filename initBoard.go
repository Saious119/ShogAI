package main

func InitBoard() [][]string { //creates a new board for a new game and prints it
	NewBoard := [][]string{
		{"L1", "N1", "S1", "G1", "K1", "G1", "S1", "N1", "L1"},
		{"O", "R1", "O", "O", "O", "O", "O", "B1", "O"},
		{"P1", "P1", "P1", "P1", "P1", "P1", "P1", "P1", "P1"},
		{"O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O"},
		{"O", "O", "O", "O", "O", "O", "O", "O"},
		{"P2", "P2", "P2", "P2", "P2", "P2", "P2", "P2", "P2"},
		{"O", "R2", "O", "O", "O", "O", "O", "B2", "O"},
		{"L2", "N2", "S2", "G2", "K2", "G2", "S2", "N2", "L2"},
	}
	//printBoard(NewBoard)
	return NewBoard
}
