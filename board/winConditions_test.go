package board

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("WinConditions", func() {

	emptyCells := [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	topRowMatch := [][]string{
		{"X", "X", "X"},
		{"O", "O", ""},
		{"", "", ""},
	}

	leftColumnMatch := [][]string{
		{"X", "O", "O"},
		{"X", "", ""},
		{"X", "", ""},
	}

	diagonalMatch := [][]string{
		{"X", "", "O"},
		{"", "X", ""},
		{"O", "", "X"},
	}

	drawnCells := [][]string{
		{"X", "X", "O"},
		{"O", "X", "X"},
		{"X", "O", "O"},
	}

	oWin := [][]string{
		{"O", "O", "O"},
		{"", "", ""},
		{"X", "X", ""},
	}

	Context("win checkers", func() {
		Context("hasRowWin", func() {
			It("returns false for an empty board", func() {
				Expect(hasRowWin(emptyCells, "X")).To(BeFalse())
				Expect(hasRowWin(emptyCells, "O")).To(BeFalse())
			})

			It("returns true for a top row matches", func() {
				Expect(hasRowWin(topRowMatch, "X")).To(BeTrue())
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
				Expect(hasRowWin(oWin, "O")).To(BeTrue())
			})

		})

		Context("hasColWin", func() {
			It("returns false if no column wins", func() {
				Expect(hasColWin(emptyCells, "X")).To(BeFalse())
				Expect(hasColWin(emptyCells, "O")).To(BeFalse())
			})

			It("returns true if first column matches", func() {
				Expect(hasColWin(leftColumnMatch, "X")).To(BeTrue())
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
				Expect(hasDiagonalWin(emptyCells, "X")).To(BeFalse())
			})

			It("returns true if matching forward diagonal", func() {
				Expect(hasDiagonalWin(diagonalMatch, "X")).To(BeTrue())
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
			Expect(HasWin(emptyCells, "X")).To(BeFalse())
		})

		It("returns true for a draw game", func() {
			Expect(HasWin(drawnCells, "O")).To(BeFalse())
		})

		It("returns true for a row win", func() {
			Expect(HasWin(topRowMatch, "X")).To(BeTrue())
		})

		It("returns true for a column win", func() {
			Expect(HasWin(leftColumnMatch, "X")).To(BeTrue())
		})

		It("returns true for a diagonal win", func() {
			Expect(HasWin(diagonalMatch, "X")).To(BeTrue())
		})

		It("returns true if O wins", func() {
			Expect(HasWin(oWin, "O")).To(BeTrue())
		})
	})
})
