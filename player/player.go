package player

import (
	"TicTacGo/board"
	"fmt"
)

type Player interface {
	PickMove(board.Board, string) (int, int)
}

func NewPlayer(playerType string) (Player, error) {
	switch playerType {
	case "human":
		return newHumanPlayer(), nil
	default:
		return nil, fmt.Errorf("invalid player type: %s", playerType)
	}
}
