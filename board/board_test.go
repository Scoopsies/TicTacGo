package board_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sut "TicTacGo/board"
)

var _ = Describe("Board", func() {

	var board sut.Board

	BeforeEach(func() {
		board = sut.CreateBoard(3)
	})

	Context("Create Board", func() {
		It("creates a board", func() {
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

	Context("Make Move", func() {
		It("plays X on [0][0] for first turn", func() {
			board.AddMove(0, 0)
			cell00 := board.GetCells()[0][0]
			Expect(cell00).To(Equal("X"))
		})

		It("plays O on second turn", func() {
			board.AddMove(0, 0)
			board.AddMove(1, 1)
			cell11 := board.GetCells()[1][1]
			Expect(cell11).To(Equal("O"))
		})
	})

	Context("GetTurn", func() {
		It("returns X on an empty board", func() {
			Expect(board.GetTurn()).To(Equal("X"))
		})

		It("returns O on a board with a lone X on it", func() {
			board.AddMove(0, 0)
			Expect(board.GetTurn()).To(Equal("O"))

			board = sut.CreateBoard(3)
			board.AddMove(1, 1)
			Expect(board.GetTurn()).To(Equal("O"))
		})

	})

})
