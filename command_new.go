package main

import (
	"bufio"
	"fmt"
	"os"
)

// Standard game sizes:
// 8 x 8 (or 9 x 9), 10 mines
// 16 x 16, 40 mines
// 16 x 30, 99
func commandNew(cfg *config, args []string) error {
	reader := bufio.NewScanner(os.Stdin)

	for {
	
		fmt.Println("Choose a level of difficulty:")
		fmt.Println()
		fmt.Println("Beginner:     9x9 grid,   10 mines")
		fmt.Println("Intermediate: 16x16 grid, 40 mines")
		fmt.Println("Expert:       16x30 grid, 99 mines")

		fmt.Println()
		fmt.Print("Difficulty Level (exit to return to current game) > ")
		reader.Scan()
		words := cleanInput(reader.Text())

		switch words[0] {
		case "beginner":
			cfg.board = NewBoard(9, 9, 10)
			cfg.state = StateStart
			return nil
		case "intermediate":
			cfg.board = NewBoard(16, 16, 40)
			cfg.state = StateStart
			return nil
		case "expert":
			cfg.board = NewBoard(16, 30, 99)
			cfg.state = StateStart
			return nil
		case "exit":
			return nil
		}
	}
}