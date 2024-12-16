package core

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetCurrentToken", func() {

	var cells [][]string

	BeforeEach(func() {
		cells = [][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}
	})

	Context("GetTurn", func() {
		It("returns X on an empty board", func() {
			Expect(GetCurrentToken(cells)).To(Equal("X"))
		})

		It("returns O on a board with a lone X on [0][0]", func() {
			cells[0][0] = "X"
			Expect(GetCurrentToken(cells)).To(Equal("O"))
		})

		It("returns 0 on a board with a lone X on [1][1]", func() {
			cells[1][1] = "X"
			Expect(GetCurrentToken(cells)).To(Equal("O"))
		})

		It("returns X on a board with equal X's and O's", func() {
			cells[1][1] = "X"
			cells[0][0] = "O"
			Expect(GetCurrentToken(cells)).To(Equal("X"))
		})
	})
})
