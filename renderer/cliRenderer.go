package renderer

import (
	"TicTacGo/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type CliRenderer struct{}

func NewCliRenderer() *CliRenderer {
	return &CliRenderer{}
}

func addCellNumbers(cells [][]string) {
	for rowIdx, row := range cells {
		for colIdx, cell := range row {
			if cell == "" {
				cells[rowIdx][colIdx] = strconv.Itoa(rowIdx*len(cells) + colIdx + 1)
			}
		}
	}
}

func cellsToString3x3(board interfaces.Board) string {
	cells := board.GetCells()
	addCellNumbers(cells)
	row1 := strings.Join(cells[0], " | ")
	row2 := strings.Join(cells[1], " | ")
	row3 := strings.Join(cells[2], " | ")

	return fmt.Sprintf(" %s \n %s \n %s ", row1, row2, row3)
}

func CellsToString(board interfaces.Board) string {
	switch board.GetType() {
	case "3x3":
		return cellsToString3x3(board)
	default:
		return ""
	}
}

func (r CliRenderer) Render(board interfaces.Board) {
	cellString := CellsToString(board)
	fmt.Println("")
	fmt.Println(cellString)
	fmt.Println("")
}

func (r CliRenderer) RenderMessage(message string) {
	fmt.Println(message + "\n")
}
