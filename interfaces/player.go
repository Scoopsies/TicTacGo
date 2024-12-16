package interfaces

type Player interface {
	PickMove(Board) (int, int)
	GetName() string
}
