package player_test

import (
	"TicTacGo/board"
	"TicTacGo/interfaces"
	sut "TicTacGo/player"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AiHard", func() {
	var (
		aiHard       *sut.AiHard
		allMoves     [][]int
		threeByThree interfaces.Board
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
		aiHard = sut.NewAiHard("Cody Ai")
		allMoves = [][]int{topLeft, topMid, topRight, midLeft, midMid, midRight, botLeft, botMid, botRight}
	})

	Context("getName", func() {
		It("returns its name", func() {
			aiHard2 := sut.NewAiHard("Scoopsinator5000")
			Expect(aiHard.GetName()).To(Equal("Cody Ai"))
			Expect(aiHard2.GetName()).To(Equal("Scoopsinator5000"))
		})
	})

	Context("PickMove", func() {
		Context("3x3 board", func() {
			BeforeEach(func() {
				threeByThree = board.NewThreeByThree()
			})
			It("picks a corner move as first move", func() {
				Expect(aiHard.PickMove(threeByThree)).To(Equal(topLeft))
			})

			It("takes a win when presented", func() {
				threeByThree.AddMove(aiHard.PickMove(threeByThree))
				threeByThree.AddMove(topMid)
				threeByThree.AddMove(aiHard.PickMove(threeByThree))
				threeByThree.AddMove(topRight)
				threeByThree.AddMove(aiHard.PickMove(threeByThree))
				Expect(threeByThree.GetState()).To(Equal("X"))
			})

			Context("picks the middle if", func() {
				It("the top left corner is taken", func() {
					threeByThree.AddMove(topLeft)
					Expect(aiHard.PickMove(threeByThree)).To(Equal(midMid))
				})

				It("the top right corner is taken", func() {
					threeByThree.AddMove(topRight)
					Expect(aiHard.PickMove(threeByThree)).To(Equal(midMid))
				})

				It("the bottom left corner is taken", func() {
					threeByThree.AddMove(botLeft)
					Expect(aiHard.PickMove(threeByThree)).To(Equal(midMid))
				})

				It("the bottom right corner is taken", func() {
					threeByThree.AddMove(botRight)
					Expect(aiHard.PickMove(threeByThree)).To(Equal(midMid))
				})
			})
		})

		It("ends in a draw when playing against its self", func() {
			for i := 0; i < 8; i++ {
				threeByThree.AddMove(aiHard.PickMove(threeByThree))
			}
			Expect(threeByThree.GetState()).To(Equal("draw"))
		})
	})

	Context("Picks a win if available", func() {
		It("on position 0,0", func() {
			board := mockBoard{availableMoves: allMoves, winPosition: []int{0, 0}, losePosition: []int{1, 1}}
			Expect(aiHard.PickMove(board)).To(Equal([]int{0, 0}))
		})

		It("on position 1,1", func() {
			board := mockBoard{availableMoves: allMoves, winPosition: []int{1, 1}, losePosition: []int{0, 0}}
			Expect(aiHard.PickMove(board)).To(Equal([]int{1, 1}))
		})
	})

	Context("Picks a block if no win available", func() {
		It("on position 0,0", func() {
			board := mockBoard{availableMoves: allMoves, winPosition: []int{-1, -1}, losePosition: []int{0, 0}}
			Expect(aiHard.PickMove(board)).To(Equal([]int{0, 0}))
		})

		It("on position 1,1", func() {
			board := mockBoard{availableMoves: allMoves, winPosition: []int{-1, -1}, losePosition: []int{1, 1}}
			Expect(aiHard.PickMove(board)).To(Equal([]int{1, 1}))
		})
	})

})
