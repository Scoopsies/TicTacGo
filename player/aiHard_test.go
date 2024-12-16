package player_test

import (
	sut "TicTacGo/player"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AiHard", func() {
	var aiHard *sut.AiHard

	BeforeEach(func() {
		aiHard = sut.NewAiHard("Cody Ai")
	})

	Context("GetName", func() {
		It("returns the name of the player", func() {
			aiHard2 := sut.NewAiHard("Computer")
			Expect(aiHard.GetName()).To(Equal("Cody Ai"))
			Expect(aiHard2.GetName()).To(Equal("Computer"))
		})
	})

	Context("PickMove", func() {
		Context("picks a win when available", func() {
			It("if 0,0 is a winning move", func() {
				winningMove := []int{0, 0}
				board := mockBoard{winRow: 0, winCol: 0, turn: "X", availableMoves: [][]int{winningMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(winningMove))
			})

			It("if 1,1 is a winning move", func() {
				winningMove := []int{1, 1}
				board := mockBoard{winRow: 1, winCol: 1, turn: "X", availableMoves: [][]int{winningMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(winningMove))
			})

			XIt("if winning move isn't first available", func() {
				winningMove := []int{0, 0}
				board := mockBoard{winRow: 0, winCol: 0, turn: "X", availableMoves: [][]int{{3, 3}, winningMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(winningMove))
			})

		})
	})

})
