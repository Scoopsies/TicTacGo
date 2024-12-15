package renderer

import (
	"bytes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("CliRenderer", func() {
	boardSize := "3x3"
	var cells [][]string

	BeforeEach(func() {
		cells = [][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}
	})

	Context("cellsToString", func() {
		It("converts empty board cells to string", func() {
			expected := " 1 | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			Expect(expected).To(Equal(cellsToString(cells, boardSize)))
		})

		It("Converts a board that's been played on to string", func() {
			expected := " X | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			cells[0][0] = "X"
			Expect(expected).To(Equal(cellsToString(cells, boardSize)))
		})
	})

	Context("Render", func() {
		It("prints an empty board", func() {
			reader, writer, err := os.Pipe()
			Expect(err).To(BeNil())

			originalStdout := os.Stdout
			os.Stdout = writer

			renderer := newCliRenderer()
			renderer.Render(cells, boardSize)

			writer.Close()
			os.Stdout = originalStdout

			var buffer bytes.Buffer
			_, err = buffer.ReadFrom(reader)
			Expect(err).To(BeNil())

			expected := cellsToString(cells, boardSize) + "\n"

			Expect(buffer.String()).To(Equal(expected))
		})
	})

})
