package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	board GameBoard
	state GameState
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		cfg.board.ShowBoard()
		if cfg.board.state == StateLose {
			fmt.Println()
			fmt.Println("Game Over")
		} else if cfg.board.state == StateWin {
			fmt.Println()
			fmt.Println("You've won!")
		}
		fmt.Println()
		fmt.Print("Minesweeper > ")
		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := listCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command, enter 'help' for instructions")
			continue
		}
		args := []string{}
		if len(words) >= 2 {
			args = words[1:]
		}
		err := command.callback(cfg, args)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func (c cliCommand) menuString() string {
	return fmt.Sprintf("%s: %s", c.name, c.description)
}

func listCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"open": {
			name:        "open",
			description: "Reveals a chosen space",
			callback:    commandOpen,
		},
		"flag": {
			name:        "flag",
			description: "Flags a space a suspected mine",
			callback:    commandFlag,
		},
		"new": {
			name:        "new",
			description: "Starts a new game",
			callback:    commandNew,
		},
		"peek": {
			name:        "peek",
			description: "Peeks at the underlying board",
			callback:    commandPeek,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit Minesweeper",
			callback:    commandExit,
		},
	}
}
