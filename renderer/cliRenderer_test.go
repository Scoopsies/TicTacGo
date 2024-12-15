package renderer

import (
	"TicTacGo/board"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CliRenderer", func() {
	Context("cellsToString", func() {
		It("converts empty board cells to string", func() {
			board, _ := board.NewBoard("3x3")
			expected := " 1 | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "

			Expect(expected).To(Equal(cellsToString(board)))
		})

		It("Converts a board that's been played on to string", func() {
			board, _ := board.NewBoard("3x3")
			expected := " X | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			board.AddMove(0, 0)
			Expect(expected).To(Equal(cellsToString(board)))
		})
	})

	Context("Render", func() {
		XIt("prints an empty board", func() {
			board, _ := board.NewBoard("3x3")
			renderer, _ := NewRenderer("cli")
			renderer.Render(board)
		})
	})

})
