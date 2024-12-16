package core

type Board interface {
	AddMove(row, column int) error
	GetCells() [][]string
	GetType() string
	GetState() string
}

type Player interface {
	PickMove(Board, string) (int, int)
	GetName() string
}

type Renderer interface {
	Render(cells [][]string, boardSize string)
	RenderMessage(message string)
}
