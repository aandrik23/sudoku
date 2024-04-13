package main

import (
	"fmt"
	"os"
)

const (
	n     = 9
	blank = '.'
)

func solveSudoku(board [][]byte) bool {
	var row, col int
	if !findEmptyLocation(board, &row, &col) {
		return true // all cells are filled
	}
	for num := byte('1'); num <= '9'; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = blank // backtrack
		}
	}
	return false
}

func findEmptyLocation(board [][]byte, row, col *int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == blank {
				*row, *col = i, j
				return true
			}
		}
	}
	return false
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < n; i++ {
		if board[row][i] == num || board[i][col] == num || board[row-row%3+i/3][col-col%3+i%3] == num {
			return false
		}
	}
	return true
}

func printBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}
	board := make([][]byte, n)
	for i, arg := range args {
		if len(arg) != n {
			fmt.Println("Error")
			return
		}
		board[i] = []byte(arg)
	}
	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
