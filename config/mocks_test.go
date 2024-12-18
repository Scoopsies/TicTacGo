package config

import "TicTacGo/interfaces"

type MockRenderer struct {
	RenderedBoards []interfaces.Board
	Messages       []string
	Called         []string
}

func (m *MockRenderer) RenderTitle() {
}

func (m *MockRenderer) RenderBoard(board interfaces.Board) {
	m.Called = append(m.Called, "RenderBoard")
	m.RenderedBoards = append(m.RenderedBoards, board)
}

func (m *MockRenderer) RenderMessage(message string) {
	m.Called = append(m.Called, "RenderMessage")
	m.Messages = append(m.Messages, message)
}

type mockInput struct {
	Inputs []string
}

func (m *mockInput) GetInput() string { // Pointer receiver
	if len(m.Inputs) > 0 {
		input := m.Inputs[0]
		m.Inputs = m.Inputs[1:] // Modify the actual Inputs slice
		return input
	}
	return ""
}
