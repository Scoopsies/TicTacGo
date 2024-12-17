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
				board := mockBoard{turn: "X", availableMoves: [][]int{winningMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(winningMove))
			})

			It("if 1,1 is a winning move", func() {
				winningMove := []int{1, 1}
				board := mockBoard{winRow: 1, winCol: 1, turn: "X", availableMoves: [][]int{winningMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(winningMove))
			})

			It("if winning move isn't first available", func() {
				winningMove := []int{0, 0}
				board := mockBoard{turn: "X", availableMoves: [][]int{{3, 3}, winningMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(winningMove))
			})

			It("picks a win over a block", func() {
				winningMove := []int{0, 0}
				board := mockBoard{blockRow: 3, blockCol: 3, turn: "X", availableMoves: [][]int{{3, 3}, winningMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(winningMove))
			})
		})

		Context("picks a block if win isn't available", func() {
			It("if block is first available move", func() {
				blockMove := []int{1, 1}
				board := mockBoard{winRow: 1, winCol: 1, turn: "X", availableMoves: [][]int{blockMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(blockMove))
			})

			It("picks a block if it is not the first available move", func() {
				blockMove := []int{0, 0}
				board := mockBoard{winRow: 1, winCol: 1, turn: "X", availableMoves: [][]int{{3, 3}, blockMove}}
				move := aiHard.PickMove(board)
				Expect(move).To(Equal(blockMove))
			})
		})

	})

})
