package board

import "fmt"

type Board interface {
	GetTurn() string
	GetCells() [][]string
	AddMove(row, column int) error
}

func NewBoard(boardType string) (Board, error) {
	switch boardType {
	case "threeByThree":
		return NewThreeByThree(), nil
	default:
		return nil, fmt.Errorf("unsupported board type: %s", boardType)

	}
}
