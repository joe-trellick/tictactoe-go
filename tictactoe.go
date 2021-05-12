package main

import (
	"errors"
	"fmt"
	"strconv"
)

type player struct {
	name  string
	piece rune
}

func printBoard(board [3][3]rune) {
	fmt.Println()
	for rowIndex, row := range board {
		if rowIndex > 0 {
			fmt.Println("━╋━╋━")
		}
		for colIndex, val := range row {
			if colIndex > 0 {
				fmt.Print("┃")
			}
			if val == 0 {
				squareIndex, _ := indicesToSquareNumber(rowIndex, colIndex)
				fmt.Print(squareIndex)
			} else {
				var boardVal string
				if val == '❌' {
					boardVal = "X"
				} else {
					boardVal = "O"
				}
				fmt.Print(boardVal)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func squareNumberToIndices(number int) (rowIndex int, colIndex int, err error) {
	if number < 1 || number > 9 {
		return 0, 0, errors.New("Number is out of range 1 - 9")
	}

	rowIndex = (number - 1) / 3
	colIndex = (number - 1) % 3
	return
}

func indicesToSquareNumber(rowIndex int, colIndex int) (squareNumber int, err error) {
	if rowIndex < 0 || rowIndex >= 3 {
		return 0, errors.New("rowIndex out of range")
	}
	if colIndex < 0 || colIndex >= 3 {
		return 0, errors.New("colIndex out of range")
	}

	return rowIndex*3 + colIndex + 1, nil
}

func main() {
	var board [3][3]rune

	player1 := player{name: "Player 1", piece: '❌'}
	player2 := player{name: "Player 2", piece: '⭕'}
	players := [2]player{player1, player2}
	turn := 0
	var hasWinner bool

	fmt.Print("\n\nWelcome to TIC ❌ TAC ⭕ TOE\n\n")

	for !hasWinner {

		printBoard(board)

		currentPlayer := players[turn%2]
		fmt.Print(string(currentPlayer.piece) + " " + currentPlayer.name + ", please enter your move: ")
		var moveInput string
		fmt.Scanln(&moveInput)

		number, inputError := strconv.Atoi(moveInput)
		rowIndex, colIndex, inputError2 := squareNumberToIndices(number)

		if inputError != nil || inputError2 != nil {
			fmt.Println("Please enter a square number from the board")
			continue
		}

		board[rowIndex][colIndex] = currentPlayer.piece
		turn += 1
	}
}
