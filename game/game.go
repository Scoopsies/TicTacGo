package game

import (
	"TicTacGo/interfaces"
)

type Game struct {
	board    interfaces.Board
	playerX  interfaces.Player
	playerO  interfaces.Player
	renderer interfaces.Renderer
}

func NewGame(board interfaces.Board, playerX, playerO interfaces.Player, renderer interfaces.Renderer) *Game {
	return &Game{
		board:    board,
		playerX:  playerX,
		playerO:  playerO,
		renderer: renderer,
	}
}

func PlayGame(game *Game) {
	board := game.board
	playerX := game.playerX
	playerO := game.playerO
	renderer := game.renderer
	gameState := board.GetState()

	for gameState == "inProgress" {
		renderer.Render(board)

		var currentPlayer interfaces.Player
		if board.GetTurn() == "X" {
			currentPlayer = playerX
		} else {
			currentPlayer = playerO
		}
		renderer.RenderMessage("It's " + currentPlayer.GetName() + "'s turn.")
		err := board.AddMove(currentPlayer.PickMove(board))

		if err != nil {
			renderer.RenderMessage(err.Error())
		}
		gameState = board.GetState()
	}

	renderer.Render(board)
}
