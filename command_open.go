package main

import (
	"errors"
)

func commandOpen(cfg *config, args []string) error {
	if cfg.board.state == StateLose || cfg.board.state == StateWin {
		return nil
	}
	if len(args) < 1 {
		return errors.New("No valid space given")
	}
	row, column, err := parseCoords(args[0])
	if err != nil {
		return err
	}

	return cfg.board.Open(row, column)
}
