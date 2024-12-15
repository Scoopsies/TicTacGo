package player

import (
	"TicTacGo/board"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func testPickMove(input string, row, column int) {
	threeByThree, _ := board.NewBoard("3x3")
	player, _ := NewPlayer("human", "Scoops")
	r, c := player.PickMove(threeByThree, input)
	Expect(r).To(Equal(row))
	Expect(c).To(Equal(column))
}

var _ = Describe("Human", func() {
	Context("PickMove", func() {
		Context("3x3 board", func() {
			It("returns 0 0 if player inputs 1", func() {
				testPickMove("1", 0, 0)
			})

			It("returns 0 1 if player inputs 2", func() {
				testPickMove("2", 0, 1)
			})

			It("returns 0 2 if player inputs 3", func() {
				testPickMove("3", 0, 2)
			})

			It("returns 1 0 if player inputs 4", func() {
				testPickMove("4", 1, 0)
			})

			It("returns 1 1 if player inputs 5", func() {
				testPickMove("5", 1, 1)
			})

			It("returns 1 2 if player inputs 6", func() {
				testPickMove("6", 1, 2)
			})

			It("returns 2 0 if player inputs 7", func() {
				testPickMove("7", 2, 0)
			})

			It("returns 2 1 if player inputs 8", func() {
				testPickMove("8", 2, 1)
			})

			It("returns 2 2 if player inputs 9", func() {
				testPickMove("9", 2, 2)
			})

			It("returns -1 -1 for any invalid input", func() {
				testPickMove("bad input", -1, -1)
			})
		})

	})

	Context("GetName", func() {
		It("returns the name of the player", func() {
			player1, _ := NewPlayer("human", "Scoops")
			player2, _ := NewPlayer("human", "Alex")
			Expect(player1.GetName()).To(Equal("Scoops"))
			Expect(player2.GetName()).To(Equal("Alex"))
		})

	})
})
