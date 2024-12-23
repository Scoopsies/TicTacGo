package factory

import (
	"TicTacGo/interfaces"
	"TicTacGo/player"
	"fmt"
)

func NewPlayer(playerType, name string, input interfaces.Input) (interfaces.Player, error) {
	switch playerType {
	case "aiHard":
		return player.NewAiHard(name), nil
	case "aiEasy":
		return player.NewAiEasy(name), nil
	case "aiMedium":
		return player.NewAiMedium(name), nil
	case "human":
		return player.NewHuman(name, input), nil

	default:
		return nil, fmt.Errorf("unsupported player type: %s", playerType)
	}
}
