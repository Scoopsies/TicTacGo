package factory_test

import (
	"TicTacGo/board"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sut "TicTacGo/factory"
)

var _ = Describe("Board", func() {
	Context("NewBoard", func() {
		It("throws an error if unsupported board type:", func() {
			_, err := sut.NewBoard("Weird size")
			Expect(err.Error()).To(ContainSubstring("unsupported board type: Weird size"))
		})

		It("creates a Three by Three board", func() {
			expectedBoard := board.NewThreeByThree()
			board, _ := sut.NewBoard("3x3")
			Expect(expectedBoard).To(Equal(board))
		})
	})
})
