package player

import "TicTacGo/interfaces"

type AiHard struct {
	name string
}

func (a AiHard) PickMove(board interfaces.Board) []int {
	return nil
}

func (a AiHard) GetName() string {
	return ""
}

func NewAiHard(name string) *AiHard {
	return &AiHard{
		name: name,
	}
}
