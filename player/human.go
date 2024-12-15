package player

type human struct {
	name string
}

func newHumanPlayer(name string) *human {
	return &human{
		name: name,
	}
}

func getPosition(moveMap map[string]Position, input string) Position {
	if position, ok := moveMap[input]; ok {
		return position
	}
	return Position{-1, -1}
}

func (h human) PickMove(_ [][]string, boardSize, input string) (int, int) {
	var moveMap map[string]Position
	switch boardSize {
	case "3x3":
		moveMap = moveMap3x3
	default:
		return -1, -1
	}

	position := getPosition(moveMap, input)
	return position.Row, position.Col
}

func (h human) GetName() string {
	name := h.name
	return name
}

type Position struct {
	Row int
	Col int
}

var moveMap3x3 = map[string]Position{
	"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
	"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
	"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
}
