package renderer

import (
	"fmt"
)

type Renderer interface {
	Render(cells [][]string, boardSize string)
	RenderMessage(message string)
}

func NewRenderer(renderType string) (Renderer, error) {
	switch renderType {
	case "cli":
		return newCliRenderer(), nil
	default:
		return nil, fmt.Errorf("unsupported render type: %s", renderType)
	}
}
