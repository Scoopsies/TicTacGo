package main

import (
	"TicTacGo/factory"
	"TicTacGo/game"
)

func main() {
	board, boardErr := factory.NewBoard("3x3")
	input, iErr := factory.NewInput("cli")
	playerX, xErr := factory.NewPlayer("human", "Scoops", input)
	playerO, oErr := factory.NewPlayer("human", "Svea", input)
	renderer, renderErr := factory.NewRenderer("cli")
	errs := []error{iErr, boardErr, xErr, oErr, renderErr}
	for _, err := range errs {
		if err != nil {
			panic(err.Error())
		}
	}
	thisGame := game.NewGame(board, playerX, playerO, renderer)

	game.PlayGame(thisGame)
}
