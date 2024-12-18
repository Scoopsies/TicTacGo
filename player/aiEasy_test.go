package player_test

import (
	sut "TicTacGo/player"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AiEasy", func() {
	var aiEasy *sut.AiEasy

	BeforeEach(func() {
		aiEasy = sut.NewAiEasy("Scoops")
	})

	Context("GetName", func() {
		It("returns its name if it's Scoops", func() {
			aiEasy := sut.NewAiEasy("Scoops")
			Expect(aiEasy.GetName()).To(Equal("Scoops"))
		})

		It("returns its name if it's Alex", func() {
			aiEasy := sut.NewAiEasy("Alex")
			Expect(aiEasy.GetName()).To(Equal("Alex"))
		})
	})

	Context("PickMove", func() {
		It("returns the only item in an array", func() {
			board := mockBoard{availableMoves: [][]int{{0, 0}}}
			move := aiEasy.PickMove(board)
			Expect(move).To(Equal([]int{0, 0}))
		})

		It("doesn't pick the same element every time", func() {
			hasNotReturnedOne := true
			hasNotReturnedTwo := true

			board := mockBoard{availableMoves: [][]int{{1}, {2}}}
			for hasNotReturnedOne || hasNotReturnedTwo {
				move := aiEasy.PickMove(board)
				if move[0] == 1 {
					hasNotReturnedOne = false
				}
				if move[0] == 2 {
					hasNotReturnedTwo = false
				}
			}

			Expect(hasNotReturnedOne).To(BeFalse())
			Expect(hasNotReturnedTwo).To(BeFalse())
		})
	})
})
