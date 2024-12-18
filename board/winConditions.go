package board

func allMatch(row []string, token string) bool {
	for _, cell := range row {
		if cell != token {
			return false
		}
	}
	return true
}

func hasRowWin(cells [][]string, token string) bool {
	for _, row := range cells {
		if allMatch(row, token) {
			return true
		}
	}
	return false
}

func CopyCells(cells [][]string) [][]string {
	newCells := make([][]string, len(cells))
	for i := range cells {
		newCells[i] = append([]string{}, cells[i]...)
	}
	return newCells
}

func rotateCells(cells [][]string) [][]string {
	numRows := len(cells)
	copiedCells := CopyCells(cells)

	for rowIdx, row := range cells {
		for colIdx, cell := range row {
			copiedCells[colIdx][numRows-1-rowIdx] = cell
		}
	}
	return copiedCells
}

func hasColWin(rotatedCells [][]string, token string) bool {
	rotatedCells = rotateCells(rotatedCells)
	return hasRowWin(rotatedCells, token)
}

func mirrorCells(cells [][]string) [][]string {
	copiedCells := CopyCells(cells)

	for rowIdx, row := range cells {
		for colIdx, cell := range row {
			copiedCells[rowIdx][len(row)-1-colIdx] = cell
		}
	}
	return copiedCells
}

func hasForwardDiagonalWin(cells [][]string, token string) bool {
	for row, rowCells := range cells {
		if rowCells[row] != token {
			return false
		}
	}
	return true
}

func hasDiagonalWin(cells [][]string, token string) bool {
	mirroredCells := mirrorCells(cells)
	return hasForwardDiagonalWin(cells, token) || hasForwardDiagonalWin(mirroredCells, token)
}

func HasWin(cells [][]string, token string) bool {
	return hasRowWin(cells, token) || hasColWin(cells, token) || hasDiagonalWin(cells, token)
}
