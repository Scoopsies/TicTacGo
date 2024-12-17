package config

import (
	"TicTacGo/factory"
	"TicTacGo/interfaces"
	"strings"
)

func getLowerInput(input interfaces.Input) string {
	return strings.ToLower(input.GetInput())
}

func getHumanPlayer(renderer interfaces.Renderer, input interfaces.Input, token string) interfaces.Player {
	renderer.RenderMessage("\nWhat is Player " + token + "'s name?")
	name := input.GetInput()
	player, _ := factory.NewPlayer("human", name, input)
	return player
}

func ConfigPlayer(renderer interfaces.Renderer, input interfaces.Input, token string) interfaces.Player {
	for {
		renderer.RenderMessage("\nWho will play as Player " + token + "?\n" + "Human or Computer?")

		switch getLowerInput(input) {
		case "human":
			return getHumanPlayer(renderer, input, token)
		case "computer":
			player, _ := factory.NewPlayer("aiHard", "Cody Ai "+token, input)
			return player
		default:
			renderer.RenderMessage("Invalid input. Please enter 'Human' or 'Computer'.")
		}
	}
}
