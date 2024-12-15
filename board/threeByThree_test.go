package board

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ThreeByThree", func() {
	threeByThreeTests()
})

func threeByThreeTests() {
	var board Board

	BeforeEach(func() {
		board, _ = NewBoard("3x3")
	})

	Context("Create Board", func() {
		It("creates a board", func() {
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
		It("returns an error if out of bounds", func() {
			err := board.AddMove(20, 20)
			Expect(err.Error()).To(ContainSubstring("out of bounds"))
		})

		It("returns an error if below bounds", func() {
			err := board.AddMove(-1, -1)
			Expect(err.Error()).To(ContainSubstring("out of bounds"))
		})

		It("returns an error if space already occupied", func() {
			board.AddMove(0, 0)
			err := board.AddMove(0, 0)
			Expect(err.Error()).To(ContainSubstring("cell already occupied"))
		})

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

		It("returns O on a board with a lone X on [0][0]", func() {
			board.AddMove(0, 0)
			Expect(board.GetTurn()).To(Equal("O"))
		})

		It("returns 0 on a board with a lone X on [1][1]", func() {
			board.AddMove(1, 1)
			Expect(board.GetTurn()).To(Equal("O"))
		})

	})

	Context("GetType", func() {
		It("returns 3x3", func() {
			Expect(board.GetType()).To(Equal("3x3"))
		})
	})
}
