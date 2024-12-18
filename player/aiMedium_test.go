package player_test

import (
	board2 "TicTacGo/board"
	sut "TicTacGo/player"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player/AiMedium", func() {
	var (
		aiMedium *sut.AiMedium
		allMoves [][]int
	)

	topLeft := []int{0, 0}
	topMid := []int{0, 1}
	topRight := []int{0, 2}
	midLeft := []int{1, 0}
	midMid := []int{1, 1}
	midRight := []int{1, 2}
	botLeft := []int{2, 0}
	botMid := []int{2, 1}
	botRight := []int{2, 2}

	BeforeEach(func() {
		aiMedium = sut.NewAiMedium("Cody Ai")
		allMoves = [][]int{topLeft, topMid, topRight, midLeft, midMid, midRight, botLeft, botMid, botRight}
	})

	Context("GetName", func() {
		It("returns the name Scoops", func() {
			aiMedium := sut.NewAiMedium("Scoops")
			Expect(aiMedium.GetName()).To(Equal("Scoops"))
		})

		It("returns the name Alex", func() {
			aiMedium := sut.NewAiMedium("Alex")
			Expect(aiMedium.GetName()).To(Equal("Alex"))
		})
	})

	Context("PickMove", func() {
		Context("takes a win when presented", func() {
			It("on position 0,0", func() {
				board := mockBoard{availableMoves: allMoves, winPosition: []int{0, 0}, losePosition: []int{1, 1}}
				Expect(aiMedium.PickMove(board)).To(Equal([]int{0, 0}))
			})

			It("on position 1,1", func() {
				board := mockBoard{availableMoves: allMoves, winPosition: []int{1, 1}, losePosition: []int{0, 0}}
				Expect(aiMedium.PickMove(board)).To(Equal([]int{1, 1}))
			})
		})
	})

	Context("Picks a block if no win available", func() {
		It("on position 0,0", func() {
			board := mockBoard{availableMoves: allMoves, winPosition: []int{-1, -1}, losePosition: []int{0, 0}}
			Expect(aiMedium.PickMove(board)).To(Equal([]int{0, 0}))
		})

		It("on position 1,1", func() {
			board := mockBoard{availableMoves: allMoves, winPosition: []int{-1, -1}, losePosition: []int{1, 1}}
			Expect(aiMedium.PickMove(board)).To(Equal([]int{1, 1}))
		})
	})

	Context("if no win or block is available it", func() {
		It("picks at random", func() {
			threeByThree := board2.NewThreeByThree()
			hasNotReturnedOne := true
			hasNotReturnedTwo := true

			for hasNotReturnedOne || hasNotReturnedTwo {
				move := aiMedium.PickMove(threeByThree)
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
