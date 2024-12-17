package factory_test

import (
	"TicTacGo/input"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sut "TicTacGo/factory"
)

var _ = Describe("Input", func() {
	Context("NewInput", func() {
		It("throws an error if unknown render type", func() {
			_, err := sut.NewInput("bad render type")
			Expect(err.Error()).To(ContainSubstring("invalid render type: bad render type"))
		})

		It("returns a cliInput", func() {
			cliInput, _ := sut.NewInput("cli")
			Expect(cliInput).To(Equal(input.NewCliInput()))
		})
	})
})
