package renderer

import (
	"bytes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	Expect(err).To(BeNil())
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore after function execution
	os.Stdout = writer

	f()

	writer.Close()

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(reader)
	Expect(err).To(BeNil())

	return buffer.String()
}

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
		It("returns empty string if invalid board size", func() {
			Expect("").To(Equal(cellsToString(cells, "invalid size")))
		})

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
			renderer := NewCliRenderer()
			output := captureOutput(func() {
				renderer.Render(cells, boardSize)
			})

			expected := cellsToString(cells, boardSize) + "\n"
			Expect(output).To(Equal(expected))
		})
	})

	Context("RenderMessage", func() {
		It("prints a message", func() {
			renderer := NewCliRenderer()
			output := captureOutput(func() {
				renderer.RenderMessage("This is a message")
			})

			Expect(output).To(Equal("This is a message\n"))
		})
	})

})
