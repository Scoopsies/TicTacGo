package main

import (
	"fmt"
)

func main() {
	if err := run("cli", "cli", "3x3"); err != nil {
		fmt.Println("Error:", err)
	}
}
