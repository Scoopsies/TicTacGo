package renderer

import (
	"TicTacGo/board"
	"fmt"
)

type Renderer interface {
	Render(board board.Board)
}

func NewRenderer(renderType string) (Renderer, error) {
	switch renderType {
	case "cli":
		return newCliRenderer(), nil
	default:
		return nil, fmt.Errorf("unsupported render type: %s", renderType)
	}
}
