package interfaces

type Renderer interface {
	RenderBoard(board Board)
	RenderMessage(message string)
	RenderTitle()
}
