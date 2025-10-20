# minesweeper-go

A CLI implementation of the classic Minesweeper game, built using Go.

# Commands:

-   Open - Reveals a space
    -- Usage: `open a1` Reveals space in row a, column 1
-   Flag - Flags a space as a suspected bomb
    -- Usage: `open a1` Reveals space in row a, column 1
-   New - Starts a new game at a specified difficulty
    -- Beginner: 9x9 grid, 10 mines
    -- Intermediate: 16x16 grid, 40 mines
    -- Expert: 16x30 grid, 99 mines
-   Help - Lists available commands
-   Exit - Exits Minesweeper

# Compilation

Standard compilation with the current Go compiler, i.e. `go build main.go`

**COMPATIBILITY NOTE:** Written and compiled using Go version 1.25.03
