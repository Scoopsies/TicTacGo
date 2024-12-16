package player

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
	Context("NewPlayer", func() {
		It("returns an error if invalid player type", func() {
			_, err := NewPlayer("invalid type", "Scoops")
			Expect(err.Error()).To(ContainSubstring("unsupported player type: invalid type"))
		})

		It("returns a human player", func() {
			player, _ := NewPlayer("human", "Scoops")
			expectedPlayer := newHumanPlayer("Scoops")
			Expect(expectedPlayer).To(Equal(player))
		})

		It("returns an aiHard player", func() {
			player, _ := NewPlayer("aiHard", "Computer")
			expectedPlayer := newAiHardPlayer("Computer")
			Expect(expectedPlayer).To(Equal(player))
		})
	})

})
