package interfaces

type Renderer interface {
	Render(cells [][]string, boardSize string)
	RenderMessage(message string)
}
