package player

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
	Context("NewPlayer", func() {
		It("returns an error if invalid player type", func() {
			_, err := NewPlayer("invalid type")
			Expect(err.Error()).To(ContainSubstring("invalid player type: invalid type"))
		})

		It("returns a human player", func() {
			player, _ := NewPlayer("human")
			expectedPlayer := newHumanPlayer()
			Expect(expectedPlayer).To(Equal(player))
		})
	})

})
