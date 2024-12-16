package game

import (
	"TicTacGo/core"
	"TicTacGo/factory"
)

type Game struct {
	board    core.Board
	playerX  core.Player
	playerO  core.Player
	renderer core.Renderer
	state    string
}

func NewGame(boardType, playerTypeX, playerXName, playerTypeO, playerOName, renderType string) (*Game, error) {
	_board, boardErr := factory.NewBoard(boardType)
	playerX, xErr := factory.NewPlayer(playerTypeX, playerXName)
	playerO, oErr := factory.NewPlayer(playerTypeO, playerOName)
	_renderer, renderErr := factory.NewRenderer(renderType)

	errs := []error{boardErr, xErr, oErr, renderErr}
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return &Game{
		board:    _board,
		playerX:  playerX,
		playerO:  playerO,
		renderer: _renderer,
		state:    "inProgress",
	}, nil
}
