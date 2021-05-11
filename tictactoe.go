package main

import "fmt"

type player struct {
	name  string
	piece rune
}

func printBoard(board [3][3]rune) {
	for rowIndex, row := range board {
		if rowIndex > 0 {
			fmt.Println("━╋━╋━")
		}
		for colIndex, val := range row {
			if colIndex > 0 {
				fmt.Print("┃")
			}
			if val == 0 {
				squareIndex := rowIndex*3 + colIndex + 1
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

func main() {
	var board [3][3]rune

	player1 := player{name: "Player 1", piece: '❌'}
	player2 := player{name: "Player 2", piece: '⭕'}
	players := [2]player{player1, player2}
	turn := 1

	board[2][1] = '❌'

	fmt.Print("\n\nWelcome to TIC ❌ TAC ⭕ TOE\n\n\n")

	printBoard(board)

	currentPlayer := players[turn%2]
	fmt.Println("Current player: " + currentPlayer.name + " " + string(currentPlayer.piece))
}
