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

func (b *Board) GetTurn() string {
	return "X"
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
