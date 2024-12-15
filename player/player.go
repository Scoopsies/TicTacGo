package player

import (
	"TicTacGo/board"
	"fmt"
)

type Player interface {
	PickMove(board.Board, string) (int, int)
	GetName() string
}

func NewPlayer(playerType, name string) (Player, error) {
	switch playerType {
	case "human":
		return newHumanPlayer(name), nil
	default:
		return nil, fmt.Errorf("unsupported player type: %s", playerType)
	}
}
