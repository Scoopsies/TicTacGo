package factory

import (
	"TicTacGo/input"
	"TicTacGo/interfaces"
	"fmt"
)

func NewInput(renderType string) (interfaces.Input, error) {
	switch renderType {
	case "cli":
		return &input.CliInput{}, nil
	}
	return nil, fmt.Errorf("invalid render type: %s", renderType)
}
