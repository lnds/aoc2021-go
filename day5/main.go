package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

type Point [2]int
type Line [2]Point

func (point Point) X() int {
	return point[0]
}

func (point Point) Y() int {
	return point[1]
}

func parsePoint(text string) Point {
	text = strings.TrimSpace(text)
	parts := strings.Split(text, ",")
	point := Point{}
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	point[0] = n
	n, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	point[1] = n
	return point
}

func parseLine(text string) Line {
	parts := strings.Split(text, "->")
	line := Line{}
	line[0] = parsePoint(parts[0])
	line[1] = parsePoint(parts[1])
	return line

}

type Board [][]int

func newBoard(lines []Line) Board {
	maxX := 0
	maxY := 0
	for _, line := range lines {
		if line[0].X() > maxX {
			maxX = line[0].X()
		}
		if line[0].Y() > maxY {
			maxY = line[0].Y()
		}
		if line[1].X() > maxX {
			maxX = line[1].X()
		}
		if line[1].Y() > maxY {
			maxY = line[1].Y()
		}
	}

	board := make(Board, maxY+1)
	for y := 0; y <= maxY; y++ {
		row := make([]int, maxX+1)
		board[y] = row
	}
	return board
}

func (board Board) Draw(line Line) Board {
	if line[0].X() == line[1].X() {
		return board.DrawVertical(line)
	} else if line[0].Y() == line[1].Y() {
		return board.DrawHorizontal(line)
	} else if abs(line[0].X()-line[1].X()) == abs(line[0].Y()-line[1].Y()) {
		return board.DrawDiagonal(line)
	}
	return board
}

func (board Board) DrawVertical(line Line) Board {
	j := line[0].X()
	if line[0].Y() > line[1].Y() {
		for i := line[1].Y(); i <= line[0].Y(); i++ {
			board[i][j]++
		}
	} else {
		for i := line[0].Y(); i <= line[1].Y(); i++ {
			board[i][j]++
		}
	}
	return board
}

func (board Board) DrawHorizontal(line Line) Board {
	i := line[0].Y()
	if line[0].X() > line[1].X() {
		for j := line[1].X(); j <= line[0].X(); j++ {
			board[i][j]++
		}
	} else {
		for j := line[0].X(); j <= line[1].X(); j++ {
			board[i][j]++
		}
	}
	return board
}

func (board Board) DrawDiagonal(line Line) Board {
	fmt.Printf("draw diagonal %v\n", line)
	d0 := line[0].X()*line[0].X() + line[0].Y()*line[0].Y()
	d1 := line[1].X()*line[1].X() + line[1].Y()*line[1].Y()
	fmt.Printf("d0 = %d, d1 = %d\n", d0, d1)
	steps := abs(line[0].X()-line[1].X()) + 1
	dx := 0
	dy := 0
	i := 0
	j := 0
	if d0 <= d1 {
		j = line[0].X()
		i = line[0].Y()

		if line[0].X() < line[1].X() {
			dx = 1
		} else {
			dx = -1
		}
		if line[0].Y() < line[1].Y() {
			dy = 1
		} else {
			dy = -1
		}
	} else {
		j = line[1].X()
		i = line[1].Y()
		dx = 0
		dy = 0
		if line[1].X() < line[0].X() {
			dx = 1
		} else {
			dx = -1
		}
		if line[1].Y() < line[0].Y() {
			dy = 1
		} else {
			dy = -1
		}
	}
	for steps > 0 {
		board[i][j]++
		i += dy
		j += dx
		steps--
	}
	return board
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []Line{}
	for scanner.Scan() {
		lines = append(lines, parseLine(scanner.Text()))
	}

	board := newBoard(lines)
	for _, line := range lines {
		board = board.Draw(line)
	}
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] >= 2 {
				count++
			}
		}
	}
	fmt.Printf("%v\n", board)
	fmt.Printf("result = %d\n", count)
}
