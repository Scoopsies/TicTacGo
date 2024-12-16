package player

import "TicTacGo/interfaces"

type AiHard struct {
	Name string
}

func NewAiHard(name string) *AiHard {
	return &AiHard{
		Name: name,
	}
}

func (a AiHard) PickMove(board interfaces.Board) []int {
	return nil
}

func (a AiHard) GetName() string {
	return a.Name
}
