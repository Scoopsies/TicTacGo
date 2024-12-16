package factory

import (
	"TicTacGo/board"
	"TicTacGo/input"
	"TicTacGo/interfaces"
	"TicTacGo/player"
	"TicTacGo/renderer"
	"fmt"
)

func NewBoard(boardType string) (interfaces.Board, error) {
	switch boardType {
	case "3x3":
		return board.NewThreeByThree(), nil
	default:
		return nil, fmt.Errorf("unsupported board type: %s", boardType)
	}
}

func NewPlayer(playerType, name string, input interfaces.Input) (interfaces.Player, error) {
	switch playerType {
	case "human":
		return player.NewHuman(name, input), nil
	case "aiHard":
		return player.NewAiHard(name), nil
	default:
		return nil, fmt.Errorf("unsupported player type: %s", playerType)
	}
}

func NewRenderer(renderType string) (interfaces.Renderer, error) {
	switch renderType {
	case "cli":
		return renderer.NewCliRenderer(), nil
	default:
		return nil, fmt.Errorf("unsupported render type: %s", renderType)
	}
}

func NewInput(renderType string) (interfaces.Input, error) {
	switch renderType {
	case "cli":
		return &input.CliInput{}, nil
	}
	return nil, fmt.Errorf("invalid render type: %s", renderType)
}
