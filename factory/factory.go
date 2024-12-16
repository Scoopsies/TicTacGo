package factory

import (
	"TicTacGo/board"
	"TicTacGo/core"
	"TicTacGo/player"
	"TicTacGo/renderer"
	"fmt"
)

func NewBoard(boardType string) (core.Board, error) {
	switch boardType {
	case "3x3":
		return board.NewThreeByThree(), nil
	default:
		return nil, fmt.Errorf("unsupported board type: %s", boardType)
	}
}

func NewPlayer(playerType, name string) (core.Player, error) {
	switch playerType {
	case "human":
		return player.NewHumanPlayer(name), nil
	default:
		return nil, fmt.Errorf("unsupported player type: %s", playerType)
	}
}

func NewRenderer(renderType string) (core.Renderer, error) {
	switch renderType {
	case "cli":
		return renderer.NewCliRenderer(), nil
	default:
		return nil, fmt.Errorf("unsupported render type: %s", renderType)
	}
}
