package player

import "TicTacGo/interfaces"

type Human struct {
	name  string
	input interfaces.Input
}

func NewHuman(name string, input interfaces.Input) *Human {
	return &Human{
		name:  name,
		input: input,
	}
}

func getPosition(moveMap map[string][]int, input string) []int {
	if position, ok := moveMap[input]; ok {
		return position
	}
	return []int{-1, -1}
}

var moveMap3x3 = map[string][]int{
	"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
	"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
	"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
}

func (h Human) PickMove(board interfaces.Board) []int {
	var moveMap map[string][]int
	switch board.GetType() {
	case "3x3":
		moveMap = moveMap3x3
	default:
		return []int{-1, -1}
	}

	input := h.input.GetInput()
	return getPosition(moveMap, input)
}

func (h Human) GetName() string {
	name := h.name
	return name
}
