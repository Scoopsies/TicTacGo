package renderer

import (
	"TicTacGo/board"
	"fmt"
	"strconv"
	"strings"
)

type cliRenderer struct{}

func addCellNumbers(cells [][]string) {
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[row]); col++ {
			if cells[row][col] == "" {
				cells[row][col] = strconv.Itoa(row*len(cells) + col + 1)
			}
		}
	}
}

func cellsToString3x3(board board.Board) string {
	cells := board.GetCells()
	addCellNumbers(cells)
	row1 := strings.Join(cells[0], " | ")
	row2 := strings.Join(cells[1], " | ")
	row3 := strings.Join(cells[2], " | ")

	return fmt.Sprintf(" %s \n %s \n %s ", row1, row2, row3)
}

func cellsToString(board board.Board) string {
	switch board.GetType() {
	case "3x3":
		return cellsToString3x3(board)
	default:
		return ""
	}
}

func (r cliRenderer) Render(board board.Board) {
	cellString := cellsToString(board)
	fmt.Println(cellString)
}

func newCliRenderer() *cliRenderer {
	return &cliRenderer{}
}
