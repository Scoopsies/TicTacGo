package board

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("WinConditions", func() {

	var board Board

	BeforeEach(func() {
		board, _ = NewBoard("3x3")
	})

	Context("win checkers", func() {
		Context("hasRowWin", func() {
			It("returns false for an empty board", func() {
				cells := board.GetCells()
				Expect(hasRowWin(cells, "X")).To(BeFalse())
				Expect(hasRowWin(cells, "O")).To(BeFalse())
			})

			It("returns true for a top row matches", func() {
				cells := [][]string{
					{"X", "X", "X"},
					{"O", "O", ""},
					{"", "", ""},
				}
				Expect(hasRowWin(cells, "X")).To(BeTrue())
			})

			It("returns true for a middle row matches", func() {
				cells := [][]string{
					{"O", "O", ""},
					{"X", "X", "X"},
					{"", "", ""},
				}
				Expect(hasRowWin(cells, "X")).To(BeTrue())
			})

			It("returns true for a bottom row matches", func() {
				cells := [][]string{
					{"O", "O", ""},
					{"", "", ""},
					{"X", "X", "X"},
				}
				Expect(hasRowWin(cells, "X")).To(BeTrue())
			})

			It("returns true if O win", func() {
				cells := [][]string{
					{"O", "O", "O"},
					{"", "", ""},
					{"X", "X", ""},
				}
				Expect(hasRowWin(cells, "O")).To(BeTrue())
			})

		})

		Context("hasColWin", func() {
			It("returns false if no column wins", func() {
				cells := board.GetCells()
				Expect(hasColWin(cells, "X")).To(BeFalse())
				Expect(hasColWin(cells, "O")).To(BeFalse())
			})

			It("returns true if first column matches", func() {
				cells := [][]string{
					{"X", "O", "O"},
					{"X", "", ""},
					{"X", "", ""},
				}
				Expect(hasColWin(cells, "X")).To(BeTrue())
			})

			It("returns true if second column matches", func() {
				cells := [][]string{
					{"O", "X", "O"},
					{"", "X", ""},
					{"", "X", ""},
				}
				Expect(hasColWin(cells, "X")).To(BeTrue())
			})

			It("returns true if third column matches", func() {
				cells := [][]string{
					{"", "O", "X"},
					{"", "", "X"},
					{"", "", "X"},
				}
				Expect(hasColWin(cells, "X")).To(BeTrue())
			})

			It("returns true if O has a matching column", func() {
				cells := [][]string{
					{"", "X", "O"},
					{"", "", "O"},
					{"", "", "O"},
				}
				Expect(hasColWin(cells, "O")).To(BeTrue())
			})
		})

		Context("hasDiagonalWin", func() {
			It("returns false if no matching diagonals", func() {
				cells := board.GetCells()
				Expect(hasDiagonalWin(cells, "X")).To(BeFalse())
			})

			It("returns true if matching forward diagonal", func() {
				cells := [][]string{
					{"X", "", "O"},
					{"", "X", ""},
					{"O", "", "X"},
				}
				Expect(hasDiagonalWin(cells, "X")).To(BeTrue())
			})

			It("returns true if matching backwards diagonal", func() {
				cells := [][]string{
					{"O", "", "X"},
					{"", "X", ""},
					{"X", "", "O"},
				}
				Expect(hasDiagonalWin(cells, "X")).To(BeTrue())
			})
		})
	})

	Context("hasWin", func() {
		It("returns false for an empty board", func() {
			cells := board.GetCells()
			Expect(hasWin(cells, "X")).To(BeFalse())
		})

		It("returns true for a row win", func() {
			board.AddMove(0, 0)
			board.AddMove(1, 0)
			board.AddMove(0, 1)
			board.AddMove(1, 1)
			board.AddMove(0, 2)
			cells := board.GetCells()
			Expect(hasWin(cells, "X")).To(BeTrue())
		})

		It("returns true for a column win", func() {
			board.AddMove(0, 0)
			board.AddMove(0, 1)
			board.AddMove(1, 0)
			board.AddMove(0, 2)
			board.AddMove(2, 0)
			cells := board.GetCells()
			Expect(hasWin(cells, "X")).To(BeTrue())
		})

		It("returns true for a diagonal win", func() {
			board.AddMove(0, 0)
			board.AddMove(0, 1)
			board.AddMove(1, 1)
			board.AddMove(0, 2)
			board.AddMove(2, 2)
			cells := board.GetCells()
			Expect(hasWin(cells, "X")).To(BeTrue())
		})

		It("returns true for a O wins", func() {
			board.AddMove(1, 0)
			board.AddMove(0, 0)
			board.AddMove(1, 1)
			board.AddMove(0, 1)
			board.AddMove(2, 2)
			board.AddMove(0, 2)
			cells := board.GetCells()
			Expect(hasWin(cells, "O")).To(BeTrue())
		})

	})
})
