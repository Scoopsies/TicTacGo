package board

import "fmt"

type ThreeByThree struct {
	size  int
	cells [][]string
}

func (b *ThreeByThree) GetCells() [][]string {
	return b.cells
}

func countXO(cells [][]string) (int, int) {
	xCount, oCount := 0, 0
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells); col++ {
			if cells[row][col] == "X" {
				xCount++
			} else if cells[row][col] == "O" {
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
	if isNotInBounds(b, row, column) {
		return fmt.Errorf("invalid move: out of bounds")
	}
	if isOccupied(b, row, column) {
		return fmt.Errorf("invalid move: cell already occupied")
	}
	b.cells[row][column] = b.GetTurn()
	return nil
}

func NewThreeByThree() *ThreeByThree {
	return &ThreeByThree{
		size:  3,
		cells: [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}},
	}
}
