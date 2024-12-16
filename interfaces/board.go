package interfaces

type Board interface {
	AddMove(row, column int) error
	GetCells() [][]string
	GetType() string
	GetState() string
	GetAvailableMoves()
}
