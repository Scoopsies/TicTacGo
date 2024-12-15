package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Renderer", func() {
	Context("NewRenderer", func() {
		It("returns an error if invalid renderer type", func() {
			_, err := NewRenderer("invalid renderer")
			Expect(err.Error()).To(ContainSubstring("unsupported render type:"))
		})

		It("creates a cliRenderer", func() {
			expectedRenderer := newCliRenderer()
			renderer, _ := NewRenderer("cli")
			Expect(expectedRenderer).To(Equal(renderer))
		})
	})
})
