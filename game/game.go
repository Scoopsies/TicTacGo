package game

import (
	"TicTacGo/board"
	"TicTacGo/player"
	"TicTacGo/renderer"
)

type Game struct {
	board    board.Board
	playerX  player.Player
	playerO  player.Player
	renderer renderer.Renderer
	state    string
}

func NewGame(boardType, playerTypeX, playerTypeO, renderType string) (*Game, error) {
	_board, boardErr := board.NewBoard(boardType)
	playerX, xErr := player.NewPlayer(playerTypeX)
	playerO, oErr := player.NewPlayer(playerTypeO)
	_renderer, renderErr := renderer.NewRenderer(renderType)

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
