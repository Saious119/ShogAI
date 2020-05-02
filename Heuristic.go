package main

import (
	"fmt"
	"strings"
)

var PieceValue = map[string]int{
	"P1":  1,
	"L1":  3,
	"N1":  3,
	"S1":  5,
	"G1":  5,
	"B1":  8,
	"R1":  9,
	"P1+": 4,
	"L1+": 6,
	"N1+": 6,
	"S1+": 6,
	"B1+": 12,
	"R1+": 13,
	"K1":  20,
	"P2":  1,
	"L2":  3,
	"N2":  3,
	"S2":  5,
	"G2":  5,
	"B2":  8,
	"R2":  9,
	"P2+": 4,
	"L2+": 6,
	"N2+": 6,
	"S2+": 6,
	"B2+": 12,
	"R2+": 13,
	"K2":  20,
}

func h(board [][]string) int {
	var p1Value = scan(board, "1")
	fmt.Println("p1value = ", p1Value)
	var p2Value = scan(board, "2")
	fmt.Println("p2value = ", p2Value)
	h := p1Value - p2Value
	return h
}

func scan(board [][]string, player string) int {
	totalValue := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if strings.Contains(board[i][j], player) {
				totalValue += PieceValue[board[i][j]]
			}
		}
	}
	return totalValue
}
