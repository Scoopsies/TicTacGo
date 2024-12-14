package board

type Board struct {
	size  int
	cells [][]string
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) GetCells() [][]string {
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

func (b *Board) GetTurn() string {
	cells := b.cells
	xCount, oCount := countXO(cells)

	if xCount > oCount {
		return "O"
	}

	return "X"
}

func (b *Board) AddMove(row, column int) {
	b.cells[row][column] = b.GetTurn()
}

func CreateBoard(size int) Board {
	return Board{
		size:  size,
		cells: createCells(size),
	}
}

func createCells(size int) [][]string {
	cells := make([][]string, size)

	for i := 0; i < size; i++ {
		cells[i] = make([]string, size)
	}
	return cells
}
