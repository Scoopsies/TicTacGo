package main

import (
	"TicTacGo/factory"
	"TicTacGo/game"
)

func main() {
	board, boardErr := factory.NewBoard("3x3")
	input, iErr := factory.NewInput("cli")
	playerO, xErr := factory.NewPlayer("human", "Scoops", input)
	playerX, oErr := factory.NewPlayer("aiHard", "Cody Ai", nil)
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
