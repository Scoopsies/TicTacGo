package player_test

type mockBoard struct {
	size string
}

func (m mockBoard) GetTurn() string { return "" }

func (m mockBoard) GetAvailableMoves() [][]int {
	return nil
}

func (m mockBoard) AddMove(_ []int) error {
	return nil
}

func (m mockBoard) GetCells() [][]string {
	return make([][]string, 0)
}

func (m mockBoard) GetType() string {
	return m.size
}

func (m mockBoard) GetState() string {
	return ""
}
