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

func rotateCells(cells [][]string) [][]string {
	numRows := len(cells)
	rotatedCells := CopyCells(cells)

	for rowIdx, row := range cells {
		for colIdx, cell := range row {
			rotatedCells[colIdx][numRows-1-rowIdx] = cell
		}
	}
	return rotatedCells
}

func hasColWin(cells [][]string, token string) bool {
	cells = rotateCells(cells)
	return hasRowWin(cells, token)
}

func mirrorCells(cells [][]string) [][]string {
	mirroredCells := CopyCells(cells)

	for rowIdx, row := range cells {
		for colIdx, cell := range row {
			mirroredCells[rowIdx][len(row)-1-colIdx] = cell
		}
	}
	return mirroredCells
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
