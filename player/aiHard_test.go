package player_test

import (
	sut "TicTacGo/player"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AiHard", func() {
	Context("GetName", func() {
		It("returns the name of the player", func() {
			player1 := sut.NewAiHard("Cody Ai")
			player2 := sut.NewAiHard("Computer")
			Expect(player1.GetName()).To(Equal("Cody Ai"))
			Expect(player2.GetName()).To(Equal("Computer"))
		})
	})

	Context("PickMove", func() {
		Context("always picks a win if available", func() {

		})
	})

})
