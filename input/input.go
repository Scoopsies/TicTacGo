package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input interface {
	GetInput() string
}

func NewInput(renderType string) (Input, error) {
	switch renderType {
	case "cli":
		return &cliInput{}, nil
	}
	return nil, fmt.Errorf("invalid render type: %s", renderType)
}

type cliInput struct{}

func (c *cliInput) GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
