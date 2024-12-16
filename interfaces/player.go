package interfaces

type Player interface {
	PickMove(Board, string) (int, int)
	GetName() string
}
