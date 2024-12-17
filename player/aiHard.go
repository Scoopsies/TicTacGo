package player

import (
	"TicTacGo/interfaces"
)

type AiHard struct {
	Name string
}

func NewAiHard(name string) *AiHard {
	return &AiHard{
		Name: name,
	}
}

func maybeGetWin(board interfaces.Board, availableMoves [][]int) []int {
	for _, move := range availableMoves {
		if board.WouldWin(move) {
			return move
		}
	}
	return nil
}

func maybeGetBlock(board interfaces.Board, availableMoves [][]int) []int {
	for _, move := range availableMoves {
		if board.WouldBlock(move) {
			return move
		}
	}
	return nil
}

func (a AiHard) PickMove(board interfaces.Board) []int {
	availableMoves := board.GetAvailableMoves()
	winMove := maybeGetWin(board, availableMoves)
	blockMove := maybeGetBlock(board, availableMoves)

	switch {
	case winMove != nil:
		return winMove
	case blockMove != nil:
		return blockMove
	default:
		return PickBestMove(board, availableMoves)
	}
}

func (a AiHard) GetName() string {
	return a.Name
}
