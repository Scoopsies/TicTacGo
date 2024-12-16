package game_test

import (
	"TicTacGo/factory"
	sut "TicTacGo/game"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	Context("newGame", func() {
		Context("creates a game that", func() {
			input, _ := factory.NewInput("cli")
			playerX, _ := factory.NewPlayer("human", "X", input)
			playerO, _ := factory.NewPlayer("human", "O", input)
			board, _ := factory.NewBoard("3x3")
			renderer, _ := factory.NewRenderer("cli")
			game := sut.NewGame(board, playerX, playerO, renderer)

			It("has a playerX", func() {
				Expect(game.PlayerX).To(Equal(playerX))
			})

			It("has a playerO", func() {
				Expect(game.PlayerO).To(Equal(playerO))
			})

			It("has a board", func() {
				Expect(game.Board).To(Equal(board))
			})

			It("has a renderer", func() {
				Expect(game.Renderer).To(Equal(renderer))
			})
		})
	})

	Context("PlayGame", func() {
		Context("PlayerX wins the game", func() { PlayGameTest("X") })
		Context("PlayerO wins the game", func() { PlayGameTest("O") })
		Context("A drawn game", func() {
			var (
				board    *MockBoard
				playerX  *MockPlayer
				playerO  *MockPlayer
				renderer *MockRenderer
				g        *sut.Game
			)
			BeforeEach(func() {
				board = &MockBoard{State: "draw", Turn: "X", Called: []string{}}
				playerX = &MockPlayer{
					name:      "Player X",
					called:    []string{},
					iteration: -1,
					positionList: [][]int{
						{0, 0},
					},
				}
				playerO = &MockPlayer{}
				renderer = &MockRenderer{}
				g = sut.NewGame(board, playerX, playerO, renderer)
				sut.PlayGame(g)
			})

			It("Announces the game was a draw.", func() {
				Expect(renderer.Called).To(ContainElement("Render"))
				Expect(renderer.Messages).To(ContainElement("No winners this time. Draw."))
			})

			It("Does not announce the winner message", func() {
				Expect(renderer.Messages).To(Not(ContainElement("Player X wins!")))
			})
		})
		Context("A drawn game", func() {
			var (
				board    *MockBoard
				playerX  *MockPlayer
				playerO  *MockPlayer
				renderer *MockRenderer
				g        *sut.Game
			)
			BeforeEach(func() {
				board = &MockBoard{State: "draw", Turn: "X", Called: []string{}}
				playerX = &MockPlayer{
					name:      "Player X",
					called:    []string{},
					iteration: -1,
					positionList: [][]int{
						{0, 0},
					},
				}
				playerO = &MockPlayer{}
				renderer = &MockRenderer{}
				g = sut.NewGame(board, playerX, playerO, renderer)
				sut.PlayGame(g)
			})

			It("Announces the game was a draw.", func() {
				Expect(renderer.Called).To(ContainElement("Render"))
				Expect(renderer.Messages).To(ContainElement("No winners this time. Draw."))
			})

			It("Does not announce the winner message", func() {
				Expect(renderer.Messages).To(Not(ContainElement("Player X wins!")))
			})
		})
	})
})
