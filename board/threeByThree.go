package board

import "fmt"

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
	return copyCells(b.cells)
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

func (b *ThreeByThree) GetTurn() string {
	cells := b.cells
	xCount, oCount := countXO(cells)
	if xCount > oCount {
		return "O"
	}
	return "X"
}

func isNotInBounds(b *ThreeByThree, row, column int) bool {
	isBelowBounds := row < 0 || column < 0
	isAboveBounds := row >= b.size || column >= b.size
	return isBelowBounds || isAboveBounds
}

func isOccupied(b *ThreeByThree, row, column int) bool {
	return b.cells[row][column] != ""
}

func (b *ThreeByThree) AddMove(row, column int) error {
	switch {
	case isNotInBounds(b, row, column):
		return fmt.Errorf("invalid move: out of bounds")
	case isOccupied(b, row, column):
		return fmt.Errorf("invalid move: cell already occupied")
	default:
		b.cells[row][column] = b.GetTurn()
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
	case hasWin(b.cells, "X"):
		return "winX"
	case hasWin(b.cells, "O"):
		return "winO"
	case hasNoValidMoves(b.cells):
		return "draw"
	default:
		return "inProgress"
	}
}
