package renderer_test

type mockBoard struct {
	size  string
	cells [][]string
}

func (m mockBoard) WouldWin(position []int) bool {
	return true
}

func (m mockBoard) GetAvailableMoves() [][]int {
	return nil
}

func (m mockBoard) GetTurn() string { return "" }

func (m mockBoard) AddMove(_ []int) error {
	return nil
}

func (m mockBoard) GetCells() [][]string {
	return m.cells
}

func (m mockBoard) GetType() string {
	return m.size
}

func (m mockBoard) GetState() string {
	return ""
}
