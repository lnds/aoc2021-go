package main

import (
	"fmt"
	"os"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])
	grid := parse(lines)
	steps := 0
	for {
		moves := grid.step()
		steps++
		if moves == 0 {
			break
		}
	}
	fmt.Println("stop after ", steps, "steps")
}

const (
	East  = '>'
	South = 'v'
	Empty = '.'
)

type Cell rune
type Grid struct {
	cells  [][]Cell
	Width  int
	Height int
}

func parse(lines []string) *Grid {
	cells := make([][]Cell, 0)
	for _, line := range lines {
		cells = append(cells, []Cell(line))
	}
	return &Grid{
		cells:  cells,
		Width:  len(cells[0]),
		Height: len(cells),
	}
}

func (grid *Grid) canMoveEast(row, col int) bool {
	if grid.cells[row][col] != East {
		return false
	}
	if col < grid.Width-1 {
		return grid.cells[row][col+1] == Empty
	} else if col == grid.Width-1 {
		return grid.cells[row][0] == Empty
	}
	return false
}

func (grid *Grid) canMoveSouth(row, col int) bool {
	if grid.cells[row][col] != South {
		return false
	}
	if row < grid.Height-1 {
		return grid.cells[row+1][col] == Empty
	} else if row == grid.Height-1 {
		return grid.cells[0][col] == Empty
	}
	return false
}

type Pos struct{ row, col int }

func (grid *Grid) moveEastHerd() int {
	moves := []Pos{}
	for i := 0; i < grid.Height; i++ {
		for j := 0; j < grid.Width; j++ {
			if grid.canMoveEast(i, j) {
				moves = append(moves, Pos{i, j})
			}
		}
	}
	for _, move := range moves {
		if move.col == grid.Width-1 {
			grid.cells[move.row][0] = East
		} else {
			grid.cells[move.row][move.col+1] = East
		}
		grid.cells[move.row][move.col] = Empty
	}
	return len(moves)
}

func (grid *Grid) moveSouthHerd() int {
	moves := []Pos{}
	for i := 0; i < grid.Height; i++ {
		for j := 0; j < grid.Width; j++ {
			if grid.canMoveSouth(i, j) {
				moves = append(moves, Pos{i, j})
			}
		}
	}
	for _, move := range moves {
		if move.row == grid.Height-1 {
			grid.cells[0][move.col] = South
		} else {
			grid.cells[move.row+1][move.col] = South
		}
		grid.cells[move.row][move.col] = Empty
	}
	return len(moves)
}

func (grid Grid) step() int {
	me := grid.moveEastHerd()
	ms := grid.moveSouthHerd()
	return me + ms
}
