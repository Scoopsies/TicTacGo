package interfaces

type Player interface {
	PickMove(Board) []int
	GetName() string
}
