package game_test

import (
	"TicTacGo/interfaces"
	"errors"
)

type MockBoard struct {
	State     string
	Turn      string
	moves     [][]int
	Called    []string
	errorMove bool
}

func (m *MockBoard) WouldWin(position []int) bool {
	return true
}

func (m *MockBoard) AddMove(position []int) error {
	m.Called = append(m.Called, "AddMove")
	m.moves = append(m.moves, position)
	if m.errorMove {
		return errors.New("invalid move")
	}
	return nil
}

func (m *MockBoard) GetCells() [][]string {
	return make([][]string, 0)
}

func (m *MockBoard) GetType() string {
	return ""
}

func (m *MockBoard) GetAvailableMoves() [][]int {
	return nil
}

func (m *MockBoard) GetState() string {
	m.Called = append(m.Called, "GetState")
	return m.State
}

func (m *MockBoard) GetTurn() string {
	m.Called = append(m.Called, "GetTurn")
	return m.Turn
}

type MockPlayer struct {
	name         string
	called       []string
	iteration    int
	positionList [][]int
}

func (m *MockPlayer) PickMove(_ interfaces.Board) []int {
	m.called = append(m.called, "PickMove")
	m.iteration++
	return m.positionList[m.iteration]
}

func (m *MockPlayer) GetName() string {
	m.called = append(m.called, "GetName")
	return m.name
}

type MockRenderer struct {
	RenderedBoards []interfaces.Board
	Messages       []string
	Called         []string
}

func (m *MockRenderer) Render(board interfaces.Board) {
	m.Called = append(m.Called, "Render")
	m.RenderedBoards = append(m.RenderedBoards, board)
}

func (m *MockRenderer) RenderMessage(message string) {
	m.Called = append(m.Called, "RenderMessage")
	m.Messages = append(m.Messages, message)
}
