package board

import (
	"TicTacGo/interfaces"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func addMoves(board interfaces.Board, moves [][]int) {
	for _, move := range moves {
		board.AddMove(move)
	}
}

var _ = Describe("ThreeByThree", func() {
	threeByThreeTests()
})

func threeByThreeTests() {
	var board interfaces.Board

	BeforeEach(func() {
		board = NewThreeByThree()
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

	Context("AddMove", func() {
		It("returns an error if out of bounds", func() {
			err := board.AddMove([]int{20, 20})
			Expect(err.Error()).To(ContainSubstring("out of bounds"))
		})

		It("returns an error if below bounds", func() {
			err := board.AddMove([]int{-1, -1})
			Expect(err.Error()).To(ContainSubstring("out of bounds"))
		})

		It("returns an error if space already occupied", func() {
			board.AddMove([]int{0, 0})
			err := board.AddMove([]int{0, 0})
			Expect(err.Error()).To(ContainSubstring("cell already occupied"))
		})

		It("plays X on first turn", func() {
			board.AddMove([]int{0, 0})
			cell00 := board.GetCells()[0][0]
			Expect(cell00).To(Equal("X"))
		})

		It("plays O on second turn", func() {
			addMoves(board, [][]int{{0, 0}, {1, 1}})
			cell11 := board.GetCells()[1][1]
			Expect(cell11).To(Equal("O"))
		})
	})

	Context("GetType", func() {
		It("returns 3x3", func() {
			Expect(board.GetType()).To(Equal("3x3"))
		})
	})

	Context("GetState", func() {
		It("returns in progress for a new game", func() {
			Expect(board.GetState()).To(Equal("inProgress"))
		})

		It("returns winX if x wins", func() {
			addMoves(board, [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}})
			Expect(board.GetState()).To(Equal("X"))
		})

		It("returns winO if o wins", func() {
			addMoves(board, [][]int{{1, 0}, {0, 0}, {1, 1}, {0, 1}, {2, 2}, {0, 2}})
			cells := board.GetCells()
			Expect(HasWin(cells, "X")).To(BeFalse())
			Expect(HasWin(cells, "O")).To(BeTrue())
			Expect(board.GetState()).To(Equal("O"))
		})

		It("returns draw if no wins and no moves left", func() {
			makeDraw(board)
			cells := board.GetCells()
			Expect(HasWin(cells, "X")).To(BeFalse())
			Expect(HasWin(cells, "O")).To(BeFalse())
			Expect(board.GetState()).To(Equal("draw"))
		})
	})

	Context("GetTurn", func() {
		It("returns X on an empty board", func() {
			Expect(board.GetTurn()).To(Equal("X"))
		})

		It("returns O on a board with a lone X on [0][0]", func() {
			board.AddMove([]int{0, 0})
			Expect(board.GetTurn()).To(Equal("O"))
		})

		It("returns 0 on a board with a lone X on [1][1]", func() {
			board.AddMove([]int{1, 1})
			Expect(board.GetTurn()).To(Equal("O"))
		})

		It("returns X on a board with equal X's and O's", func() {
			board.AddMove([]int{1, 1})
			board.AddMove([]int{0, 0})
			Expect(board.GetTurn()).To(Equal("X"))
		})
	})

	Context("getAvailableMoves", func() {
		It("returns empty []int if no moves available", func() {
			makeDraw(board)
			Expect(board.GetAvailableMoves()).To(BeNil())
		})

		It("returns all positions on empty board", func() {
			allMoves := [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
			Expect(board.GetAvailableMoves()).To(Equal(allMoves))
		})

		It("returns all positions on a partially filled board", func() {
			allMoves := [][]int{{0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
			board.AddMove([]int{0, 0})
			Expect(board.GetAvailableMoves()).To(Equal(allMoves))
		})
	})

	Context("WouldWin()", func() {
		It("returns false if sent invalid position", func() {
			makeDraw(board)
			Expect(board.WouldWin([]int{-1, -1})).To(BeFalse())
		})

		It("returns false if position is already occupied", func() {
			position := []int{0, 1}
			board.AddMove(position)
			Expect(board.WouldWin(position)).To(BeFalse())
		})

		It("returns true if it would result in a win for X", func() {
			addMoves(board, [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}})
			Expect(board.WouldWin([]int{0, 2})).To(BeTrue())
		})

	})

	Context("WouldBlock", func() {
		It("returns false if out of bounds", func() {
			Expect(board.WouldBlock([]int{-1, -1})).To(BeFalse())
		})

		It("returns false if position is already occupied", func() {
			move := []int{0, 0}
			board.AddMove(move)
			Expect(board.WouldBlock(move)).To(BeFalse())
		})

		It("returns false if not a block available", func() {
			Expect(board.WouldBlock([]int{0, 0})).To(BeFalse())
		})

		It("returns true if it would block a win", func() {
			addMoves(board, [][]int{{0, 0}, {1, 0}, {0, 1}})
			Expect(board.WouldBlock([]int{0, 2})).To(BeTrue())
		})
	})

	Context("Copy", func() {
		It("returns an identical board", func() {
			boardCopy := board.Copy()
			Expect(board).To(Equal(boardCopy))
		})

		It("making changes to the copy doesn't make changes to the original", func() {
			boardCopy := board.Copy()
			boardCopy.AddMove([]int{0, 0})
			Expect(board.GetAvailableMoves()).To(Not(Equal(boardCopy.GetAvailableMoves())))
		})
	})
}

func makeDraw(board interfaces.Board) {
	addMoves(board, [][]int{{0, 0}, {0, 2}, {0, 1}, {1, 0}, {1, 1}, {2, 1}, {1, 2}, {2, 2}, {2, 0}})
}
