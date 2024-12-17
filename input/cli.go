package input

import (
	"bufio"
	"os"
	"strings"
)

type CliInput struct{}

func NewCliInput() *CliInput {
	return &CliInput{}
}

func (c *CliInput) GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
