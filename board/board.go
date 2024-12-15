package board

import "fmt"

type Board interface {
	AddMove(row, column int) error
	GetCells() [][]string
	GetTurn() string
	GetType() string
	GetState() string
}

func NewBoard(boardType string) (Board, error) {
	switch boardType {
	case "3x3":
		return NewThreeByThree(), nil
	default:
		return nil, fmt.Errorf("unsupported board type: %s", boardType)
	}
}

func copyCells(cells [][]string) [][]string {
	newCells := make([][]string, len(cells))
	for i := range cells {
		newCells[i] = append([]string{}, cells[i]...)
	}
	return newCells
}
