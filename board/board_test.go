package board

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	Context("NewBoard", func() {
		It("throws an error if unsupported board type:", func() {
			_, err := NewBoard("Weird size")
			Expect(err.Error()).To(ContainSubstring("unsupported board type: Weird size"))
		})

		It("creates a Three by Three board", func() {
			expectedBoard := NewThreeByThree()
			board, _ := NewBoard("3x3")
			Expect(expectedBoard).To(Equal(board))
		})
	})

})
