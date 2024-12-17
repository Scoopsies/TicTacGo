package player

import "TicTacGo/interfaces"

const (
	MaxScore = 1000
	MinScore = -1000
)

func isEmpty(moves [][]int) bool {
	return len(moves) == 0
}

func Minimax(board interfaces.Board, depth int, isMaximizing bool) int {
	availableMoves := board.GetAvailableMoves()

	for _, move := range availableMoves {
		if board.WouldWin(move) {
			if isMaximizing {
				return 10 - depth
			} else {
				return depth - 10
			}
		}
	}

	if isEmpty(availableMoves) {
		return 0
	}

	if isMaximizing {
		bestScore := MinScore
		for _, move := range availableMoves {
			boardCopy := board.Copy()
			boardCopy.AddMove(move)
			score := Minimax(boardCopy, depth+1, false)
			bestScore = max(bestScore, score)
		}
		return bestScore
	} else {
		bestScore := MaxScore
		for _, move := range availableMoves {
			boardCopy := board.Copy()
			boardCopy.AddMove(move)
			score := Minimax(boardCopy, depth+1, true)
			bestScore = min(bestScore, score)
		}
		return bestScore
	}
}

func PickBestMove(board interfaces.Board, availableMoves [][]int) []int {
	bestScore := -1000
	var bestMove []int

	for _, move := range availableMoves {
		boardCopy := board.Copy()
		boardCopy.AddMove(move)

		score := Minimax(boardCopy, 0, false)

		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}
