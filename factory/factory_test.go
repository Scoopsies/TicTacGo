package factory

import (
	"TicTacGo/board"
	"TicTacGo/input"
	"TicTacGo/player"
	"TicTacGo/renderer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	Context("NewBoard", func() {
		It("throws an error if unsupported board type:", func() {
			_, err := NewBoard("Weird size")
			Expect(err.Error()).To(ContainSubstring("unsupported board type: Weird size"))
		})

		It("creates a Three by Three board", func() {
			expectedBoard := board.NewThreeByThree()
			board, _ := NewBoard("3x3")
			Expect(expectedBoard).To(Equal(board))
		})
	})

	Context("NewPlayer", func() {
		It("returns an error if invalid player type", func() {
			_, err := NewPlayer("invalid type", "Scoops")
			Expect(err.Error()).To(ContainSubstring("unsupported player type: invalid type"))
		})

		It("returns a human player", func() {
			human, _ := NewPlayer("human", "Scoops")
			expectedPlayer := player.NewHumanPlayer("Scoops")
			Expect(expectedPlayer).To(Equal(human))
		})
	})

	Context("NewRenderer", func() {
		It("returns an error if invalid renderer type", func() {
			_, err := NewRenderer("invalid renderer")
			Expect(err.Error()).To(ContainSubstring("unsupported render type:"))
		})

		It("creates a cliRenderer", func() {
			expectedRenderer := renderer.NewCliRenderer()
			cliRenderer, _ := NewRenderer("cli")
			Expect(expectedRenderer).To(Equal(cliRenderer))
		})
	})

	Context("NewInput", func() {
		It("throws an error if unknown render type", func() {
			_, err := NewInput("bad render type")
			Expect(err.Error()).To(ContainSubstring("invalid render type: bad render type"))
		})

		It("returns a cliInput", func() {
			cliInput, _ := NewInput("cli")
			Expect(cliInput).To(Equal(input.NewCliInput()))
		})
	})
})
