package player

import "TicTacGo/interfaces"

type AiMedium struct {
	name string
}

func (a AiMedium) PickMove(board interfaces.Board) []int {
	availableMoves := board.GetAvailableMoves()
	winMove := maybeGetWin(board, availableMoves)
	blockMove := maybeGetBlock(board, availableMoves)

	switch {
	case winMove != nil:
		return winMove
	case blockMove != nil:
		return blockMove
	default:
		return GetRandomPosition(board)
	}
}

func (a AiMedium) GetName() string {
	return a.name
}

func NewAiMedium(name string) *AiMedium {
	return &AiMedium{name: name}
}
