package main

import (
	"TicTacGo/factory"
	"fmt"
)

func main() {
	rendererer, _ := factory.NewRenderer("cli")
	rendererer.RenderTitle()
	if err := run("cli", "cli", "3x3"); err != nil {
		fmt.Println("Error:", err)
	}
}
