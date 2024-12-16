package game

import (
	"TicTacGo/interfaces"
)

type Game struct {
	Board    interfaces.Board
	PlayerX  interfaces.Player
	PlayerO  interfaces.Player
	Renderer interfaces.Renderer
}

func NewGame(board interfaces.Board, playerX, playerO interfaces.Player, renderer interfaces.Renderer) *Game {
	return &Game{
		Board:    board,
		PlayerX:  playerX,
		PlayerO:  playerO,
		Renderer: renderer,
	}
}

func getCurrentPlayer(board interfaces.Board, playerX, playerO interfaces.Player) interfaces.Player {
	if board.GetTurn() == "X" {
		return playerX
	}
	return playerO
}

func makeMove(board interfaces.Board, player interfaces.Player) error {
	playerMove := player.PickMove(board)
	return board.AddMove(playerMove)
}

func maybeRenderError(err error, renderer interfaces.Renderer) {
	if err != nil {
		renderer.RenderMessage(err.Error())
	}
}

func renderMovePrompt(renderer interfaces.Renderer, player interfaces.Player, board interfaces.Board) {
	renderer.Render(board)
	renderer.RenderMessage("It's " + player.GetName() + "'s turn.")
}

func renderEndScreen(renderer interfaces.Renderer, playerX, playerO interfaces.Player, board interfaces.Board) {
	renderer.Render(board)
	switch board.GetState() {
	case "draw":
		renderer.RenderMessage("No winners this time. Draw.")
	case "X":
		renderer.RenderMessage(playerX.GetName() + " wins!")
	case "O":
		renderer.RenderMessage(playerO.GetName() + " wins!")
	}
}

func loopGameplay(gameState string, board interfaces.Board, playerX, playerO interfaces.Player, renderer interfaces.Renderer) {
	for gameState == "inProgress" {
		currentPlayer := getCurrentPlayer(board, playerX, playerO)
		renderMovePrompt(renderer, currentPlayer, board)
		err := makeMove(board, currentPlayer)
		maybeRenderError(err, renderer)
		gameState = board.GetState()
	}
}

func PlayGame(game *Game) {
	board := game.Board
	playerX := game.PlayerX
	playerO := game.PlayerO
	renderer := game.Renderer
	gameState := "inProgress"

	loopGameplay(gameState, board, playerX, playerO, renderer)
	renderEndScreen(renderer, playerX, playerO, board)
}
