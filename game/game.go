package game

import (
	"TicTacGo/factory"
	"TicTacGo/interfaces"
)

type Game struct {
	board    interfaces.Board
	playerX  interfaces.Player
	playerO  interfaces.Player
	renderer interfaces.Renderer
	state    string
}

func NewGame(boardType, playerTypeX, playerXName, playerTypeO, playerOName, renderType string) (*Game, error) {
	board, boardErr := factory.NewBoard(boardType)
	playerX, xErr := factory.NewPlayer(playerTypeX, playerXName)
	playerO, oErr := factory.NewPlayer(playerTypeO, playerOName)
	renderer, renderErr := factory.NewRenderer(renderType)

	errs := []error{boardErr, xErr, oErr, renderErr}
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return &Game{
		board:    board,
		playerX:  playerX,
		playerO:  playerO,
		renderer: renderer,
		state:    "inProgress",
	}, nil
}
