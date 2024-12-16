package factory_test

import (
	"TicTacGo/board"
	sut "TicTacGo/factory"
	"TicTacGo/input"
	"TicTacGo/player"
	"TicTacGo/renderer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	Context("NewBoard", func() {
		It("throws an error if unsupported board type:", func() {
			_, err := sut.NewBoard("Weird size")
			Expect(err.Error()).To(ContainSubstring("unsupported board type: Weird size"))
		})

		It("creates a Three by Three board", func() {
			expectedBoard := board.NewThreeByThree()
			board, _ := sut.NewBoard("3x3")
			Expect(expectedBoard).To(Equal(board))
		})
	})

	Context("NewInput", func() {
		It("throws an error if unknown render type", func() {
			_, err := sut.NewInput("bad render type")
			Expect(err.Error()).To(ContainSubstring("invalid render type: bad render type"))
		})

		It("returns a cliInput", func() {
			cliInput, _ := sut.NewInput("cli")
			Expect(cliInput).To(Equal(input.NewCliInput()))
		})
	})

	Context("NewPlayer", func() {
		input, _ := sut.NewInput("cli")

		It("returns an error if invalid player type", func() {
			_, err := sut.NewPlayer("invalid type", "Scoops", input)
			Expect(err.Error()).To(ContainSubstring("unsupported player type: invalid type"))
		})

		It("returns a human player", func() {
			human, _ := sut.NewPlayer("human", "Scoops", input)
			expectedPlayer := player.NewHuman("Scoops", input)
			Expect(expectedPlayer).To(Equal(human))
		})

		It("returns an AiHard player", func() {
			aiHard, _ := sut.NewPlayer("aiHard", "Cody Ai", nil)
			expectedPlayer := player.NewAiHard("Cody Ai")
			Expect(expectedPlayer).To(Equal(aiHard))
		})
	})

	Context("NewRenderer", func() {
		It("returns an error if invalid renderer type", func() {
			_, err := sut.NewRenderer("invalid renderer")
			Expect(err.Error()).To(ContainSubstring("unsupported render type:"))
		})

		It("creates a cliRenderer", func() {
			expectedRenderer := renderer.NewCliRenderer()
			cliRenderer, _ := sut.NewRenderer("cli")
			Expect(expectedRenderer).To(Equal(cliRenderer))
		})
	})

})
