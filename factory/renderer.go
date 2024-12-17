package factory

import (
	"TicTacGo/interfaces"
	"TicTacGo/renderer"
	"fmt"
)

func NewRenderer(renderType string) (interfaces.Renderer, error) {
	switch renderType {
	case "cli":
		return renderer.NewCliRenderer(), nil
	default:
		return nil, fmt.Errorf("unsupported render type: %s", renderType)
	}
}
