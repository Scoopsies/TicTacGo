package board

import (
	"TicTacGo/interfaces"
	"fmt"
)

type ThreeByThree struct {
	size  int
	cells [][]string
}

func getOppToken(board interfaces.Board) string {
	if board.GetTurn() == "X" {
		return "O"
	} else {
		return "X"
	}
}

func (b *ThreeByThree) WouldBlock(position []int) bool {
	row, col := position[0], position[1]
	if isNotInBounds(b, row, col) || isOccupied(b, row, col) {
		return false
	}

	token := getOppToken(b)
	cells := b.GetCells()
	cells[row][col] = token
	return HasWin(cells, token)
}

func (b *ThreeByThree) WouldWin(position []int) bool {
	row, col := position[0], position[1]
	if isNotInBounds(b, row, col) || isOccupied(b, row, col) {
		return false
	}

	token := b.GetTurn()
	cells := b.GetCells()
	cells[row][col] = token
	return HasWin(cells, token)
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

func (b *ThreeByThree) GetAvailableMoves() [][]int {
	var result [][]int

	for rowIdx, row := range b.cells {
		for colIdx, col := range row {
			if col == "" {
				result = append(result, []int{rowIdx, colIdx})
			}
		}
	}

	return result
}

func GetCurrentToken(cells [][]string) string {
	xCount, oCount := countXO(cells)
	if xCount > oCount {
		return "O"
	}
	return "X"
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
	return GetCurrentToken(b.cells)
}

func isNotInBounds(b *ThreeByThree, row, column int) bool {
	isBelowBounds := row < 0 || column < 0
	isAboveBounds := row >= b.size || column >= b.size
	return isBelowBounds || isAboveBounds
}

func isOccupied(b *ThreeByThree, row, column int) bool {
	return b.cells[row][column] != ""
}

func (b *ThreeByThree) AddMove(position []int) error {
	row := position[0]
	column := position[1]
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
