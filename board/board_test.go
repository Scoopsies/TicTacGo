package board_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sut "TicTacGo/board"
)

var _ = Describe("Board", func() {

	Context("Create Board", func() {
		It("creates a board", func() {
			board := sut.CreateBoard(3)
			By("Where the size is 3", func() {
				Expect(board.GetSize()).To(Equal(3))
			})

			expectedCells := [][]string{
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
			}

			By("the cells are a 3x3 slice", func() {
				Expect(board.GetCells()).To(Equal(expectedCells))
			})
		})
	})

	Context("GetTurn", func() {
		It("returns X on an empty board", func() {
			board := sut.CreateBoard(3)
			Expect(board.GetTurn()).To(Equal("X"))
		})
	})

})
