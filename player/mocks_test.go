package player_test

import "fmt"

type mockBoard struct {
	size           string
	winRow         int
	winCol         int
	blockRow       int
	blockCol       int
	isWin          bool
	availableMoves [][]int
	turn           string
}

func (m mockBoard) WouldBlock(position []int) bool {
	row, col := position[0], position[1]
	return row == m.blockRow && col == m.blockCol
}

func (m mockBoard) WouldWin(position []int) bool {
	row, col := position[0], position[1]
	return row == m.winRow && col == m.winCol
}

func (m mockBoard) GetTurn() string { return m.turn }

func (m mockBoard) GetAvailableMoves() [][]int {
	return m.availableMoves
}

func (m mockBoard) AddMove(move []int) error {
	if move[0] == m.winRow && move[1] == m.winCol {
		m.isWin = true
		fmt.Println("")
		fmt.Print(move)
		fmt.Println(" triggered a win")
	}
	return nil
}

func (m mockBoard) GetCells() [][]string {
	return make([][]string, 0)
}

func (m mockBoard) GetType() string {
	return m.size
}

func (m mockBoard) GetState() string {
	if m.isWin {
		fmt.Println("dropped a token")
		return m.turn
	}
	return "O"
}
