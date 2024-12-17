package player_test

import (
	"TicTacGo/interfaces"
)

type mockBoard struct {
	size           string
	winPosition    []int
	losePosition   []int
	isWin          bool
	availableMoves [][]int
	turn           string
}

func (m mockBoard) Copy() interfaces.Board {
	return mockBoard{
		m.size, m.winPosition, m.losePosition,
		m.isWin, m.availableMoves, m.turn,
	}
}

func isEqual(position1, position2 []int) bool {
	row, col := position1[0], position1[1]
	lRow, lCol := position2[0], position2[1]
	return row == lRow && col == lCol
}

func (m mockBoard) WouldBlock(position []int) bool {
	return isEqual(position, m.losePosition)
}

func (m mockBoard) WouldWin(position []int) bool {
	return isEqual(position, m.winPosition)
}

func (m mockBoard) GetTurn() string { return m.turn }

func (m mockBoard) GetAvailableMoves() [][]int {
	return m.availableMoves
}

func (m mockBoard) AddMove(_ []int) error {
	m.availableMoves = m.availableMoves[1:]
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
		return m.turn
	}
	return "O"
}
