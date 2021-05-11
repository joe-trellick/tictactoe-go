package main

import "fmt"

type player struct {
	name  string
	piece rune
}

func main() {
	player1 := player{name: "Player 1", piece: '❌'}
	player2 := player{name: "Player 2", piece: '⭕'}
	players := [2]player{player1, player2}
	turn := 0

	fmt.Print("\n\nWelcome to TIC ❌ TAC ⭕ TOE\n\n\n")

	currentPlayer := players[turn%1]
	fmt.Println("Current player: " + currentPlayer.name + " " + string(currentPlayer.piece))
}
