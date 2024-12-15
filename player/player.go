package player

import (
	"fmt"
)

type Player interface {
	PickMove([][]string, string, string) (int, int)
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
