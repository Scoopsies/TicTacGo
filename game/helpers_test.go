package game_test

import (
	"TicTacGo/game"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func PlayGameTest(playerToken string) {
	var (
		board         *MockBoard
		playerX       *MockPlayer
		playerO       *MockPlayer
		renderer      *MockRenderer
		g             game.Game
		currentPlayer *MockPlayer
	)
	BeforeEach(func() {
		board = &MockBoard{State: playerToken, Turn: playerToken, Called: []string{}}
		playerX = &MockPlayer{
			name:      "Player X",
			called:    []string{},
			iteration: -1,
			positionList: [][]int{
				{0, 0},
			},
		}
		playerO = &MockPlayer{
			name:      "Player O",
			called:    []string{},
			iteration: -1,
			positionList: [][]int{
				{0, 0},
			},
		}
		renderer = &MockRenderer{}
		g = game.NewGame(renderer, playerX, playerO, board)
		g.Play()

		playerMap := map[string]*MockPlayer{
			"X": playerX,
			"O": playerO,
		}
		currentPlayer = playerMap[playerToken]

	})

	It("should not render an error message if invalid move", func() {
		Expect(renderer.Messages).To(Not(ContainElement("invalid move")))
	})

	It("should render the game board", func() {
		Expect(renderer.Called).To(ContainElement("RenderBoard"))
	})

	It("should prompt the current player to make move", func() {
		Expect(renderer.Called).To(ContainElement("RenderMessage"))
		Expect(renderer.Messages).To(ContainElement("It's Player " + playerToken + "'s turn."))
	})

	It("should invoke player.PickMove", func() {
		Expect(currentPlayer.called).To(ContainElement("PickMove"))
	})

	It("should invoke board.AddMove with result of player.PickMove", func() {
		Expect(board.Called).To(ContainElement("AddMove"))
		Expect(board.moves).To(ContainElement(currentPlayer.positionList[0]))
	})

	It("announces the winner", func() {
		Expect(renderer.Called).To(ContainElement("RenderBoard"))
		Expect(renderer.Messages).To(ContainElement("Player " + playerToken + " wins!"))
	})

	It("does not announce a draw", func() {
		Expect(renderer.Messages).To(Not(ContainElement("No winners this time. Draw.")))
	})
}
