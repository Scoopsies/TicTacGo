package board

func CopyCells(cells [][]string) [][]string {
	newCells := make([][]string, len(cells))
	for i := range cells {
		newCells[i] = append([]string{}, cells[i]...)
	}
	return newCells
}
