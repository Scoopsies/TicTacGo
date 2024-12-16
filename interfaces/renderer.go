package interfaces

type Renderer interface {
	Render(board Board)
	RenderMessage(message string)
}
