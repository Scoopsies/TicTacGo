package renderer_test

import (
	sut "TicTacGo/renderer"
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
			board := mockBoard{size: "invalid board size"}
			Expect("").To(Equal(sut.CellsToString(board)))
		})

		It("converts empty board cells to string", func() {
			expected := " 1 | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			board := mockBoard{cells: cells, size: "3x3"}
			Expect(expected).To(Equal(sut.CellsToString(board)))
		})

		It("Converts a board that's been played on to string", func() {
			expected := " X | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			cells[0][0] = "X"
			board := mockBoard{cells: cells, size: "3x3"}
			Expect(expected).To(Equal(sut.CellsToString(board)))
		})
	})

	Context("RenderBoard", func() {
		It("prints an empty board", func() {
			renderer := sut.NewCliRenderer()
			board := mockBoard{
				cells: cells,
				size:  "3x3",
			}

			output := captureOutput(func() {
				renderer.RenderBoard(board)
			})

			expected := "\n" + sut.CellsToString(board) + "\n\n"
			Expect(output).To(Equal(expected))
		})
	})

	Context("RenderMessage", func() {
		It("prints a message", func() {
			renderer := sut.NewCliRenderer()
			output := captureOutput(func() {
				renderer.RenderMessage("This is a message")
			})

			Expect(output).To(Equal("This is a message\n\n"))
		})

		It("prints a different message", func() {
			renderer := sut.NewCliRenderer()
			output := captureOutput(func() {
				renderer.RenderMessage("This is a different message")
			})

			Expect(output).To(Equal("This is a different message\n\n"))
		})
	})

	Context("RenderTitle", func() {
		It("renders the title", func() {
			renderer := sut.NewCliRenderer()
			output := captureOutput(func() {
				renderer.RenderTitle()
			})
			Expect(output).To(Equal(sut.CliTitle + "\n"))
		})
	})

})
