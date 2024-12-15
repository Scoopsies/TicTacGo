package renderer

import (
	"TicTacGo/board"
	"bytes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("CliRenderer", func() {

	var board3x3 board.Board

	BeforeEach(func() {
		board3x3, _ = board.NewBoard("3x3")
	})

	Context("cellsToString", func() {
		It("converts empty board cells to string", func() {
			expected := " 1 | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			Expect(expected).To(Equal(cellsToString(board3x3)))
		})

		It("Converts a board that's been played on to string", func() {
			expected := " X | 2 | 3 \n" +
				" 4 | 5 | 6 \n" +
				" 7 | 8 | 9 "
			board3x3.AddMove(0, 0)
			Expect(expected).To(Equal(cellsToString(board3x3)))
		})
	})

	Context("Render", func() {
		It("prints an empty board", func() {
			reader, writer, err := os.Pipe()
			Expect(err).To(BeNil())

			originalStdout := os.Stdout
			os.Stdout = writer

			renderer := newCliRenderer()
			renderer.Render(board3x3)

			writer.Close()
			os.Stdout = originalStdout

			var buffer bytes.Buffer
			_, err = buffer.ReadFrom(reader)
			Expect(err).To(BeNil())

			expected := cellsToString(board3x3) + "\n"

			Expect(buffer.String()).To(Equal(expected))
		})
	})

})
