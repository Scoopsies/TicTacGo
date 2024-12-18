package player

import (
	"TicTacGo/interfaces"
	"math/rand"
)

type AiEasy struct {
	name string
}

func NewAiEasy(name string) *AiEasy {
	return &AiEasy{name}
}

func GetRandomPosition(board interfaces.Board) []int {
	moves := board.GetAvailableMoves()
	return moves[rand.Intn(len(moves))]
}

func (a AiEasy) PickMove(board interfaces.Board) []int {
	return GetRandomPosition(board)
}

func (a AiEasy) GetName() string {
	return a.name
}
