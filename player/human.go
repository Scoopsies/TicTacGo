package player

import "TicTacGo/board"

type human struct {
}

func (h human) PickMove(_ board.Board) {

}

func newHumanPlayer() *human {
	return &human{}
}
