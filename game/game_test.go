package game

import (
	"TicTacGo/board"
	"TicTacGo/player"
	"TicTacGo/renderer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {

	Context("newGame", func() {
		Context("creates a game that", func() {

			game, err := NewGame("3x3", "human", "X", "human", "O", "cli")
			playerX, _ := player.NewPlayer("human", "X")
			playerO, _ := player.NewPlayer("human", "O")
			board, _ := board.NewBoard("3x3")
			_renderer, _ := renderer.NewRenderer("cli")

			It("has a playerX", func() {
				Expect(game.playerX).To(Equal(playerX))
			})

			It("has a playerO", func() {
				Expect(game.playerO).To(Equal(playerO))
			})

			It("has a state starting with inProgress", func() {
				Expect(game.state).To(Equal("inProgress"))
			})

			It("has a board", func() {
				Expect(game.board).To(Equal(board))
			})

			It("has a renderer", func() {
				Expect(game.renderer).To(Equal(_renderer))
			})

			It("does not send an error if all valid input", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("has an error if", func() {
			It("boardType is invalid", func() {
				_, err := NewGame("bad board", "human", "X", "human", "O", "cli")
				Expect(err.Error()).To(Equal("unsupported board type: bad board"))
			})

			It("playerTypeX is invalid", func() {
				_, err := NewGame("3x3", "bad player", "X", "human", "O", "cli")
				Expect(err.Error()).To(Equal("unsupported player type: bad player"))
			})

			It("playerTypeO is invalid", func() {
				_, err := NewGame("3x3", "human", "X", "bad player", "O", "cli")
				Expect(err.Error()).To(Equal("unsupported player type: bad player"))
			})

			It("renderType is invalid", func() {
				_, err := NewGame("3x3", "human", "X", "human", "O", "bad renderer")
				Expect(err.Error()).To(Equal("unsupported render type: bad renderer"))
			})
		})

		Context("PlayGame", func() {

		})

	})
})
