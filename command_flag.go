package main

import (
	"errors"
)

func commandFlag(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("No valid space given")
	}
	row, column, err := parseCoords(args[0])
	if err != nil {
		return err
	}
	return cfg.gameBoard.Flag(row, column)
}