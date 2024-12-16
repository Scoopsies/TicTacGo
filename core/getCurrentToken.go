package core

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
