package renderer

import (
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

func cellsToString3x3(cells [][]string) string {
	addCellNumbers(cells)
	row1 := strings.Join(cells[0], " | ")
	row2 := strings.Join(cells[1], " | ")
	row3 := strings.Join(cells[2], " | ")

	return fmt.Sprintf(" %s \n %s \n %s ", row1, row2, row3)
}

func cellsToString(cells [][]string, boardSize string) string {
	switch boardSize {
	case "3x3":
		return cellsToString3x3(cells)
	default:
		return ""
	}
}

func (r CliRenderer) Render(cells [][]string, boardSize string) {
	cellString := cellsToString(cells, boardSize)
	fmt.Println(cellString)
}

func (r CliRenderer) RenderMessage(message string) {
	fmt.Println(message)
}
