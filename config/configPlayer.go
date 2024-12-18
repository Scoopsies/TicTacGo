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
	player, _ := factory.NewPlayer("human", input.GetInput(), input)
	return player
}

func capitalize(input string) string {
	if len(input) == 0 {
		return ""
	}
	return strings.ToUpper(string(input[0])) + input[1:]
}

func getComputerPlayer(renderer interfaces.Renderer, input interfaces.Input, token string) interfaces.Player {
	for {
		renderer.RenderMessage("\nWhat difficulty for Computer " + token + "?\nEasy, Medium, or Hard")

		selection := getLowerInput(input)
		aiType := "ai" + capitalize(selection)
		name := "Cody Ai " + token
		player, err := factory.NewPlayer(aiType, name, input)

		if err != nil {
			renderer.RenderMessage(selection + " is not a valid difficulty.")
			continue
		}

		return player
	}
}

func ConfigPlayer(renderer interfaces.Renderer, input interfaces.Input, token string) interfaces.Player {
	for {
		renderer.RenderMessage("\nWho will play as Player " + token + "?\nHuman or Computer?")

		switch getLowerInput(input) {
		case "human":
			return getHumanPlayer(renderer, input, token)
		case "computer":
			return getComputerPlayer(renderer, input, token)
		default:
			renderer.RenderMessage("Invalid input. Please enter 'Human' or 'Computer'.")
		}
	}
}
