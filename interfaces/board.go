package interfaces

type Board interface {
	AddMove(position []int) error
	GetCells() [][]string
	GetType() string
	GetState() string
	GetAvailableMoves() [][]int
	GetTurn() string
}
