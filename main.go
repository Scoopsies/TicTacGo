package main

import (
	"TicTacGo/config"
	"TicTacGo/factory"
	game "TicTacGo/game"
)

func maybePanic(errors []error) {
	for _, err := range errors {
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	renderer, rErr := factory.NewRenderer("cli")
	input, iErr := factory.NewInput("cli")
	playerX := config.ConfigPlayer(renderer, input, "X")
	playerO := config.ConfigPlayer(renderer, input, "O")
	board, bErr := factory.NewBoard("3x3")

	maybePanic([]error{rErr, iErr, bErr})

	ticTacToe := game.NewGame(renderer, playerX, playerO, board)
	ticTacToe.Play()
}
