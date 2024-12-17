package factory_test

import (
	"TicTacGo/renderer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sut "TicTacGo/factory"
)

var _ = Describe("Renderer", func() {
	Context("NewRenderer", func() {
		It("returns an error if invalid renderer type", func() {
			_, err := sut.NewRenderer("invalid renderer")
			Expect(err.Error()).To(ContainSubstring("unsupported render type:"))
		})

		It("creates a cliRenderer", func() {
			expectedRenderer := renderer.NewCliRenderer()
			cliRenderer, _ := sut.NewRenderer("cli")
			Expect(expectedRenderer).To(Equal(cliRenderer))
		})
	})
})
