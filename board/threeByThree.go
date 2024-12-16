package board

import (
	"fmt"
)

type ThreeByThree struct {
	size  int
	cells [][]string
}

func NewThreeByThree() *ThreeByThree {
	return &ThreeByThree{
		size:  3,
		cells: [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}},
	}
}

func (b *ThreeByThree) GetCells() [][]string {

	return CopyCells(b.cells)
}

func isNotInBounds(b *ThreeByThree, row, column int) bool {
	isBelowBounds := row < 0 || column < 0
	isAboveBounds := row >= b.size || column >= b.size
	return isBelowBounds || isAboveBounds
}

func isOccupied(b *ThreeByThree, row, column int) bool {
	return b.cells[row][column] != ""
}

func countXO(cells [][]string) (int, int) {
	xCount, oCount := 0, 0
	for _, row := range cells {
		for _, cell := range row {
			switch cell {
			case "X":
				xCount++
			case "O":
				oCount++
			}
		}
	}
	return xCount, oCount
}

func GetCurrentToken(cells [][]string) string {
	xCount, oCount := countXO(cells)
	if xCount > oCount {
		return "O"
	}
	return "X"
}

func (b *ThreeByThree) AddMove(row, column int) error {
	switch {
	case isNotInBounds(b, row, column):
		return fmt.Errorf("invalid move: out of bounds")
	case isOccupied(b, row, column):
		return fmt.Errorf("invalid move: cell already occupied")
	default:
		b.cells[row][column] = GetCurrentToken(b.cells)
		return nil
	}
}

func (b *ThreeByThree) GetType() string {
	return "3x3"
}

func hasNoValidMoves(cells [][]string) bool {
	for _, row := range cells {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func (b *ThreeByThree) GetState() string {
	switch {
	case HasWin(b.cells, "X"):
		return "X"
	case HasWin(b.cells, "O"):
		return "O"
	case hasNoValidMoves(b.cells):
		return "draw"
	default:
		return "inProgress"
	}
}

func (b *ThreeByThree) GetAvailableMoves() {

}
