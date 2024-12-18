package factory_test

import (
	"TicTacGo/player"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sut "TicTacGo/factory"
)

var _ = Describe("Player", func() {
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

		It("returns an AiEasy player", func() {
			aiEasy, _ := sut.NewPlayer("aiEasy", "Cody Ai", nil)
			expectedPlayer := player.NewAiEasy("Cody Ai")
			Expect(expectedPlayer).To(Equal(aiEasy))
		})
	})
})
