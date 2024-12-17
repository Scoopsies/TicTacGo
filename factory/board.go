package factory

import (
	"TicTacGo/board"
	"TicTacGo/interfaces"
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
