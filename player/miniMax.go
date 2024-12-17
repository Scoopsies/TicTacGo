package player

import "TicTacGo/interfaces"

const (
	MaxScore  = 1000
	MinScore  = -1000
	WinScore  = 10
	DrawScore = 0
)

func isEmpty(moves [][]int) bool {
	return len(moves) == 0
}

func hasWinningMove(board interfaces.Board, moves [][]int) bool {
	for _, move := range moves {
		if board.WouldWin(move) {
			return true
		}
	}
	return false
}

func scoreWin(depth int, isMaximizing bool) int {
	if isMaximizing {
		return WinScore - depth
	} else {
		return depth - WinScore
	}
}

func Minimax(board interfaces.Board, depth int, isMaximizing bool) int {
	availableMoves := board.GetAvailableMoves()

	switch {
	case isEmpty(availableMoves):
		return DrawScore
	case hasWinningMove(board, availableMoves):
		return scoreWin(depth, isMaximizing)
	case isMaximizing:
		bestScore := MinScore
		for _, move := range availableMoves {
			boardCopy := board.Copy()
			boardCopy.AddMove(move)
			score := Minimax(boardCopy, depth+1, false)
			bestScore = max(bestScore, score)
		}
		return bestScore
	default:
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
