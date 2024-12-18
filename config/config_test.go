package config

import (
	"TicTacGo/factory"
	"TicTacGo/interfaces"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func expectPromptToContain(renderer *MockRenderer, s string, input interfaces.Input, token string) {
	runConfigPlayer(renderer, input, token)
	Expect(renderer.Messages).To(ContainElement(s))
}

func runConfigPlayer(renderer *MockRenderer, input interfaces.Input, token string) interfaces.Player {
	return ConfigPlayer(renderer, input, token)
}

var _ = Describe("Config", func() {
	var renderer *MockRenderer

	BeforeEach(func() {
		renderer = &MockRenderer{}
	})

	Context("ConfigBoard", func() {
		It("prompts the user to see who plays as X", func() {
			expectPromptToContain(renderer, "\nWho will play as Player X?\nHuman or Computer?", &mockInput{[]string{"human", "Scoops"}}, "X")
		})

		It("prompts a user to see who plays as O", func() {
			expectPromptToContain(renderer, "\nWho will play as Player O?\nHuman or Computer?", &mockInput{[]string{"human", "Scoops"}}, "O")
		})

		It("prompts the user for a name if they select human for X", func() {
			expectPromptToContain(renderer, "\nWhat is Player X's name?", &mockInput{[]string{"human", "Scoops"}}, "X")
		})

		It("prompts the user for a name if they select human for O", func() {
			expectPromptToContain(renderer, "\nWhat is Player O's name?", &mockInput{[]string{"human", "Scoops"}}, "O")
		})

		It("it does not prompt for name if user selects computer", func() {
			runConfigPlayer(renderer, &mockInput{[]string{"Computer", "Hard"}}, "X")
			Expect(renderer.Messages).To(Not(ContainElement("What is Player X's name?")))
		})

		It("it prompts difficulty selection if player selects computer", func() {
			runConfigPlayer(renderer, &mockInput{[]string{"Computer", "Hard"}}, "X")
			Expect(renderer.Messages).To(ContainElement("\nWhat difficulty for Computer X?\nEasy, Medium, or Hard"))
		})

		Context("returns a human player", func() {

			It("named Scoops", func() {
				inputMock := &mockInput{[]string{"human", "Scoops"}}
				player := runConfigPlayer(renderer, inputMock, "X")
				human, _ := factory.NewPlayer("human", "Scoops", inputMock)
				Expect(player).To(Equal(human))
			})

			It("named Alex", func() {
				inputMock := &mockInput{[]string{"human", "Alex"}}
				human, _ := factory.NewPlayer("human", "Alex", inputMock)
				player := runConfigPlayer(renderer, inputMock, "X")
				Expect(player).To(Equal(human))
			})

			It("if human is capitalized in input", func() {
				inputMock := &mockInput{[]string{"Human", "Scoops"}}
				human, _ := factory.NewPlayer("human", "Scoops", inputMock)
				player := runConfigPlayer(renderer, inputMock, "X")
				Expect(player).To(Equal(human))
			})
		})

		Context("returns a computer player if user selects computer", func() {
			It("named Cody Ai X for X", func() {
				computer, _ := factory.NewPlayer("aiHard", "Cody Ai X", nil)
				player := runConfigPlayer(renderer, &mockInput{[]string{"computer", "Hard"}}, "X")
				Expect(player).To(Equal(computer))
			})

			It("named Cody Ai O for O", func() {
				computer, _ := factory.NewPlayer("aiHard", "Cody Ai O", nil)
				player := runConfigPlayer(renderer, &mockInput{[]string{"computer", "Hard"}}, "O")
				Expect(player).To(Equal(computer))
			})

			It("returns a aiMedium if user selects medium", func() {
				computer, _ := factory.NewPlayer("aiMedium", "Cody Ai O", nil)
				player := runConfigPlayer(renderer, &mockInput{[]string{"computer", "Medium"}}, "O")
				Expect(player).To(Equal(computer))
			})

			It("returns a aiEasy if user selects easy", func() {
				computer, _ := factory.NewPlayer("aiEasy", "Cody Ai X", nil)
				player := runConfigPlayer(renderer, &mockInput{[]string{"computer", "Easy"}}, "X")
				Expect(player).To(Equal(computer))
			})

			It("prompts user if given invalid input, and then asks them to choose again", func() {
				computer, _ := factory.NewPlayer("aiEasy", "Cody Ai X", nil)
				player := runConfigPlayer(renderer, &mockInput{[]string{"computer", "bad input", "Easy"}}, "X")
				Expect(renderer.Messages).To(ContainElement("bad input is not a valid difficulty."))
				Expect(player).To(Equal(computer))
			})

			It("even if computer is capitalized.", func() {
				computer, _ := factory.NewPlayer("aiHard", "Cody Ai O", nil)
				player := runConfigPlayer(renderer, &mockInput{[]string{"COMPUTER", "HARD"}}, "O")
				Expect(player).To(Equal(computer))
			})
		})

		It("prompts user again if they don't select computer or human", func() {
			inputMock := &mockInput{[]string{"bad data", "Human", "Scoops"}}
			runConfigPlayer(renderer, inputMock, "X")
			Expect(renderer.Messages).To(ContainElement("Invalid input. Please enter 'Human' or 'Computer'."))
		})
	})

})
