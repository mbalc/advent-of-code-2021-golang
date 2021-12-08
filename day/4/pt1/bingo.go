package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type bingoField struct {
	value    int
	isMarked bool
}

type bingoBoard = [5][5]bingoField

func readBoards() []bingoBoard {
	var boards []bingoBoard
	for {
		var board bingoBoard
		_, err := fmt.Scanln(&board[0][0].value, &board[0][1].value, &board[0][2].value, &board[0][3].value, &board[0][4].value)
		// fmt.Println(n, err, board)
		if err == io.EOF {
			return boards
		}
		fmt.Scanln(&board[1][0].value, &board[1][1].value, &board[1][2].value, &board[1][3].value, &board[1][4].value)
		fmt.Scanln(&board[2][0].value, &board[2][1].value, &board[2][2].value, &board[2][3].value, &board[2][4].value)
		fmt.Scanln(&board[3][0].value, &board[3][1].value, &board[3][2].value, &board[3][3].value, &board[3][4].value)
		fmt.Scanln(&board[4][0].value, &board[4][1].value, &board[4][2].value, &board[4][3].value, &board[4][4].value)
		// fmt.Println(board)

		boards = append(boards, board)
		// fmt.Println(board)

		_, err2 := fmt.Scanln()
		if err2 == io.EOF {
			return boards
		}
	}
}

func crossOutNumber(board bingoBoard, numberToCrossOut int) bingoBoard {
	for i, line := range board {
		for j, _ := range line {
			if board[i][j].value == numberToCrossOut {
				board[i][j].isMarked = true
			}
		}
	}
	return board
}

func main() {
	var buffer string
	var inputs []string
	fmt.Scanln(&buffer)
	// fmt.Println(inputs, buffer)
	inputs = strings.Split(buffer, ",")
	// fmt.Println(inputs, buffer)

	fmt.Scanln()

	var boards []bingoBoard = readBoards()

	for _, input := range inputs {
		for i, board := range boards {
			integerInput, _ := strconv.Atoi(input)
			boards[i] = crossOutNumber(board, integerInput)
			if checkIfBingo(boards[i]) {
				unmarkedSum := calculateBoardUmarkedFieldsSum(boards[i])
				// fmt.Println(board, unmarkedSum, integerInput)
				fmt.Println(unmarkedSum * integerInput)
				return
			}
		}
	}
}

func calculateBoardUmarkedFieldsSum(board bingoBoard) int {
	var sum = 0

	for i, line := range board {
		for j, _ := range line {
			if !board[i][j].isMarked {
				sum += board[i][j].value
			}
		}
	}

	return sum
}

func checkIfBingo(board bingoBoard) bool {
	return true &&
		// rows
		(board[0][0].isMarked && board[0][1].isMarked && board[0][2].isMarked && board[0][3].isMarked && board[0][4].isMarked) ||
		(board[1][0].isMarked && board[1][1].isMarked && board[1][2].isMarked && board[1][3].isMarked && board[1][4].isMarked) ||
		(board[2][0].isMarked && board[2][1].isMarked && board[2][2].isMarked && board[2][3].isMarked && board[2][4].isMarked) ||
		(board[3][0].isMarked && board[3][1].isMarked && board[3][2].isMarked && board[3][3].isMarked && board[3][4].isMarked) ||
		(board[4][0].isMarked && board[4][1].isMarked && board[4][2].isMarked && board[4][3].isMarked && board[4][4].isMarked) ||

		// columns
		(board[0][0].isMarked && board[1][0].isMarked && board[2][0].isMarked && board[3][0].isMarked && board[4][0].isMarked) ||
		(board[0][1].isMarked && board[1][1].isMarked && board[2][1].isMarked && board[3][1].isMarked && board[4][1].isMarked) ||
		(board[0][2].isMarked && board[1][2].isMarked && board[2][2].isMarked && board[3][2].isMarked && board[4][2].isMarked) ||
		(board[0][3].isMarked && board[1][3].isMarked && board[2][3].isMarked && board[3][3].isMarked && board[4][3].isMarked) ||
		(board[0][4].isMarked && board[1][4].isMarked && board[2][4].isMarked && board[3][4].isMarked && board[4][4].isMarked)

	// diagonals
	// (board[0][0].isMarked && board[1][1].isMarked && board[2][2].isMarked && board[3][3].isMarked && board[4][4].isMarked) ||
	// (board[0][4].isMarked && board[1][3].isMarked && board[2][2].isMarked && board[3][1].isMarked && board[4][0].isMarked)
}
