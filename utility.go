package main

import (
	"strconv"
	"strings"
)

type GameState int

const (
	StateStart GameState = iota
	StateContinue
	StateWin
	StateLose
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func parseCoords(coords string) (int, int, error) {
	row := strings.Index(alphabet, strings.ToUpper(coords[:1]))
	col, err := strconv.Atoi(coords[1:])
	if err != nil {
		return 0, 0, err
	}
	col -= 1
	return row, col, nil
}
