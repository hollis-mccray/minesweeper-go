package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

type cell struct {
	is_mine       bool
	is_flagged    bool
	is_open       bool
	adjacentMines int
}

func (c cell) String() string {
	if !c.is_open {
		if c.is_flagged {
			return "⚑"
		}
		return "■"
	}

	if c.is_mine {
		return "X"
	}

	if c.adjacentMines == 0 {
		return " "
	}

	return strconv.Itoa(c.adjacentMines)
}

func (c cell) peek() string {

	if c.is_mine {
		return "X"
	}

	if c.adjacentMines == 0 {
		return " "
	}

	return strconv.Itoa(c.adjacentMines)
}

type GameBoard struct {
	cells       [][]cell
	rows, columns, mines int
	state       GameState
}

// Has to be before the factory method NewBoard because NewBoard calls it
func (b *GameBoard)placeMine(){
	for {
		row := rand.Intn(b.rows)
		col := rand.Intn(b.rows)
		if !b.cells[row][col].is_mine {
			b.cells[row][col].is_mine = true
			break
		}
	}

}

// Factory method that initializes a new board, including placing bombs
func NewBoard(rows, columns, mines int) GameBoard {
	newBoard := GameBoard{}

	newBoard.rows = rows
	newBoard.columns = columns
	newBoard.mines = mines
	newBoard.state = StateStart

	// The cells are stored as a list of lists of cells
	newBoard.cells = make([][]cell, rows)
	for i := 0; i < rows; i++ {
		newBoard.cells[i] = make([]cell, columns) // For each row, create a slice of 'cols' length
	}

	// Placing mines
	for i:= 0; i < newBoard.mines; i++{
		newBoard.placeMine()
	}

	return newBoard
}

// Peek function that displays the hidden board
func (b GameBoard)peek() {
	header := "  "
	for col := 1; col <= b.columns; col ++ {
		header = header + fmt.Sprintf(" %2d", col)
	}
	fmt.Println(header)
	fmt.Printf("   ┌─%s┐\n", strings.Repeat("───", b.columns - 1))
	for i:= 0; i < b.rows; i++ {
		row := make([]string, len(b.cells[i]))
		for j:= 0; j < b.columns; j++ {
			row[j] = b.cells[i][j].peek()
		}
		fmt.Printf("%2c │%s│\n", alphabet[i], strings.Join(row, "  "))
	}
	fmt.Printf("   └─%s┘\n", strings.Repeat("───", b.columns - 1))
}

func (b GameBoard)ShowBoard() {
	header := "  "
	for col := 1; col <= b.columns; col ++ {
		header = header + fmt.Sprintf(" %2d", col)
	}
	fmt.Println(header)
	fmt.Printf("   ┌─%s┐\n", strings.Repeat("───", b.columns - 1))
	for i:= 0; i < b.rows; i++ {
		row := make([]string, len(b.cells[i]))
		for j:= 0; j < b.columns; j++ {
			row[j] = b.cells[i][j].String()
		}
		fmt.Printf("%2c │%s│\n", alphabet[i], strings.Join(row, "  "))
	}
	fmt.Printf("   └─%s┘\n", strings.Repeat("───", b.columns - 1))

}

type index struct{
	row int
	column int
}

var shiftList = []index{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

func (b *GameBoard)countAdjacentMines() {

	count := func(row, col int) {
		counter := 0
		for _, shift := range shiftList{
			adjRow := row + shift.row
			adjCol := col + shift.column

			if adjRow < 0 || adjRow >= b.rows || adjCol < 0 || adjCol >= b.columns {
				continue
			}

			if b.cells[adjRow][adjCol].is_mine {
				counter += 1
			}
		}
		b.cells[row][col].adjacentMines = counter
	}

	var wg sync.WaitGroup

	for i:= 0; i < b.rows; i++ {
		for j:= 0; j < b.columns; j++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				count(i, j)
			}()
		}
	}

	wg.Wait()
}


func (b *GameBoard)victoryCheck() {
	var wg sync.WaitGroup

	allMinesFlagged := true
	allSafeOpened := true

	for i:= 0; i < b.rows; i++ {
		for j:= 0; j < b.columns; j++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				if b.cells[i][j].is_mine {
					if !b.cells[i][j].is_flagged {
						allMinesFlagged = false
					} 
				} else {
					if !b.cells[i][j].is_open {
						allSafeOpened = false
					}
				}
			}()
		}
	}

	wg.Wait()
	if allMinesFlagged || allSafeOpened {
		b.state = StateWin
	}
}

func (b *GameBoard)Open(row, column int) error {

	if row < 0 || row >= b.rows {
		return errors.New("invalid space")
	}

	if column < 0 || column >= b.columns {
		return errors.New("invalid space")
	}
	
	if b.cells[row][column].is_open {
		return nil
	}

	//The mercy rule - the first cell you open cannot be a mine
	// So if it is, we just place another mine and mark this space safe
	if b.state == StateStart {
		if b.cells[row][column].is_mine {
			b.placeMine()
			b.cells[row][column].is_mine = false
		}
		b.state = StateContinue
		b.countAdjacentMines()
	}

	b.cells[row][column].is_open = true

	if b.cells[row][column].is_mine {
		b.state = StateLose
	}  else if b.cells[row][column].adjacentMines == 0 {
		for _, shift := range shiftList{
			adjRow := row + shift.row
			adjCol := column + shift.column

			if adjRow < 0 || adjRow >= b.rows || adjCol < 0 || adjCol >= b.columns {
				continue
			}

			if !b.cells[adjRow][adjCol].is_open {
				b.Open(adjRow, adjCol)
			}
		}
	}
	b.victoryCheck()
	return nil
}

func (b *GameBoard)Flag(row, column int) error {
	if b.state == StateLose || b.state == StateWin {
		return nil
	}

	if row < 0 || row >= b.rows {
		return errors.New("invalid space")
	}

	if column < 0 || column >= b.columns {
		return errors.New("invalid space")
	}
	
	if b.cells[row][column].is_open {
		return nil
	}

	b.cells[row][column].is_flagged = !b.cells[row][column].is_flagged
	b.victoryCheck()
	return nil
}