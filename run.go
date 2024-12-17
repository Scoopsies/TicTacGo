package main

import (
	"TicTacGo/config"
	"TicTacGo/factory"
	"TicTacGo/game"
)

func run(rendererType, inputType, boardType string) error {
	renderer, rErr := factory.NewRenderer(rendererType)
	if rErr != nil {
		return rErr
	}

	input, iErr := factory.NewInput(inputType)
	if iErr != nil {
		return iErr
	}

	board, bErr := factory.NewBoard(boardType)
	if bErr != nil {
		return bErr
	}

	playerX := config.ConfigPlayer(renderer, input, "X")
	playerO := config.ConfigPlayer(renderer, input, "O")
	ticTacToe := game.NewGame(renderer, playerX, playerO, board)
	ticTacToe.Play()
	return nil
}
