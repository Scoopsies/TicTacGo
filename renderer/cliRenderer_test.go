package renderer

import (
	"bytes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

type mockBoard struct {
	cells [][]string
	size  string
}

func (m mockBoard) GetTurn() string {
	return ""
}

func (m mockBoard) AddMove(_ []int) error {
	return nil
}

func (m mockBoard) GetCells() [][]string {
	return m.cells
}

func (m mockBoard) GetType() string {
	return m.size
}

func (m mockBoard) GetState() string {
	return ""
}

func (m mockBoard) GetAvailableMoves() {
}

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
			Expect("").To(Equal(cellsToString(board)))
		})

		It("converts empty board cells to string", func() {
			expected := " 1 | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			board := mockBoard{cells: cells, size: "3x3"}
			Expect(expected).To(Equal(cellsToString(board)))
		})

		It("Converts a board that's been played on to string", func() {
			expected := " X | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			cells[0][0] = "X"
			board := mockBoard{cells: cells, size: "3x3"}
			Expect(expected).To(Equal(cellsToString(board)))
		})
	})

	Context("Render", func() {
		It("prints an empty board", func() {
			renderer := NewCliRenderer()
			board := mockBoard{
				cells: cells,
				size:  "3x3",
			}

			output := captureOutput(func() {
				renderer.Render(board)
			})

			expected := cellsToString(board) + "\n"
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
