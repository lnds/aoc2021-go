package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board [][]int

func newBoard(lines []string) Board {
	board := Board{}
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		row := []int{}
		for _, num := range numbers {
			n, e := strconv.Atoi(num)
			if e == nil {
				row = append(row, n)
			}
		}
		board = append(board, row)
	}
	return board
}

func (board Board) mark(draw int) {
	for i, row := range board {
		for j, num := range row {
			if num == draw {
				board[i][j] = -1
			}
		}
	}
}

func (board Board) rows() [][]int {
	result := [][]int{}
	for _, row := range board {
		result = append(result, row)
	}
	return result
}

func (board Board) cols() [][]int {
	result := [][]int{}
	for j := 0; j < len(board[0]); j++ {
		col := []int{}
		for i := 0; i < len(board); i++ {
			col = append(col, board[i][j])
		}
		result = append(result, col)
	}
	return result
}

func (board Board) won() bool {
	rows := board.rows()
	for _, row := range rows {
		if win(row) {
			return true
		}
	}
	cols := board.cols()
	for _, col := range cols {
		if win(col) {
			return true
		}
	}
	return false
}

func (board Board) score() int {
	sum := 0
	for _, row := range board {
		for _, cell := range row {
			if cell > 0 {
				sum += cell
			}
		}
	}
	return sum
}

func win(vec []int) bool {
	for _, c := range vec {
		if c > 0 {
			return false
		}
	}
	return true
}

func extractDraws(line string) []int {
	result := []int{}
	for _, n := range strings.Split(line, ",") {
		num, err := strconv.Atoi(n)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func playBingo(draws []int, boards []Board) int {
	for _, draw := range draws {
		for _, board := range boards {
			board.mark(draw)
			if board.won() {
				fmt.Printf("Winner board = %v\n", board)
				return board.score() * draw
			}
		}
	}
	return 0
}

func playBingo2(draws []int, boards []Board) int {
	result := 0
	winnerBoards := make([]int, len(boards))
out:
	for _, draw := range draws {
		for i, board := range boards {
			board.mark(draw)
			if board.won() {
				winnerBoards[i] = 1
				result = board.score() * draw
				winners := 0
				for _, won := range winnerBoards {
					winners += won
				}
				if winners == len(boards) {
					break out
				}
			}
		}

	}
	return result
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	draws := extractDraws(line)

	boards := []Board{}
	scanner.Scan()
	boardLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			board := newBoard(boardLines)
			boards = append(boards, board)
			boardLines = []string{}
		} else {
			boardLines = append(boardLines, line)
		}
	}
	if len(boardLines) > 0 {
		board := newBoard(boardLines)
		boards = append(boards, board)
	}
	boards2 := []Board{}
	boards2 = append(boards2, boards...)

	result := playBingo(draws, boards)
	fmt.Printf("Result 1 = %d\n", result)

	result = playBingo2(draws, boards2)
	fmt.Printf("Result 2 = %d\n", result)
}
