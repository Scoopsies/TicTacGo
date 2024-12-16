package game

import (
	"TicTacGo/factory"
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
			game := NewGame(board, playerX, playerO, renderer)

			It("has a playerX", func() {
				Expect(game.playerX).To(Equal(playerX))
			})

			It("has a playerO", func() {
				Expect(game.playerO).To(Equal(playerO))
			})

			It("has a board", func() {
				Expect(game.board).To(Equal(board))
			})

			It("has a renderer", func() {
				Expect(game.renderer).To(Equal(renderer))
			})
		})

		Context("PlayGame", func() {
			Expect("Creates")
		})

	})
})
