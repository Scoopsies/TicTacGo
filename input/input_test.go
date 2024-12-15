package input

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("input", func() {
	Context("NewInput", func() {
		It("throws an error if unknown render type", func() {
			_, err := NewInput("bad render type")
			Expect(err.Error()).To(ContainSubstring("invalid render type: bad render type"))
		})

		It("returns a cliInput", func() {
			input, _ := NewInput("cli")
			Expect(input).To(Equal(&cliInput{}))
		})
	})

	Context("cliInput", func() {
		var (
			input         *cliInput
			originalStdin *os.File
		)

		BeforeEach(func() {
			input = &cliInput{}
			originalStdin = os.Stdin
		})

		AfterEach(func() {
			os.Stdin = originalStdin
		})

		It("should return user input after trimming spaces", func() {
			readPipe, writePipe, err := os.Pipe()
			Expect(err).NotTo(HaveOccurred())
			os.Stdin = readPipe

			go func() {
				defer writePipe.Close()
				writePipe.WriteString("   tic tac go   \n")
			}()

			input := input.GetInput()
			Expect(input).To(Equal("tic tac go"))
		})

		It("should handle empty input", func() {
			readPipe, writePipe, err := os.Pipe()
			Expect(err).NotTo(HaveOccurred())
			os.Stdin = readPipe

			go func() {
				defer writePipe.Close()
				writePipe.WriteString("\n")
			}()

			input := input.GetInput()
			Expect(input).To(Equal(""))
		})
	})
})
