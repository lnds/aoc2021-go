package main

import (
	"fmt"
	"os"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])

	matrix := shared.ParseMatrix09(lines)

	flashes := 0
	for i := 1; i <= 100; i++ {
		m, f, _ := step(matrix)
		flashes += f
		matrix = m

	}
	fmt.Printf("matrix %v \nflashes = %d\n", matrix, flashes)

	all := false
	at := 0
	matrix = shared.ParseMatrix09(lines)
	for i := 1; !all; i++ {
		m, _, a := step(matrix)
		matrix = m
		all = a
		at = i
	}
	fmt.Printf("all flashed at step %d\n", at)

}

func step(matrix shared.Matrix09) (shared.Matrix09, int, bool) {
	flashed := make([][]bool, len(matrix))
	for i := range flashed {
		flashed[i] = make([]bool, len(matrix[0]))
	}

	for i, row := range matrix {
		for j := range row {
			matrix[i][j]++
		}
	}

	var directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	flashes := 0
	countFlashed := 0
	repeat := true
	for repeat {
		highEnergy := 0
		for i, row := range matrix {
			for j := range row {
				if matrix[i][j] > 9 && !flashed[i][j] {
					flashes++
					flashed[i][j] = true
					for _, dir := range directions {
						ii := i + dir[0]
						jj := j + dir[1]
						if ii >= 0 && ii < len(matrix) && jj >= 0 && jj < len(matrix[0]) {
							matrix[ii][jj]++
							if matrix[ii][jj] >= 9 && !flashed[ii][jj] {
								highEnergy++
							}
						}
					}
				}
			}
		}

		repeat = highEnergy > 0
	}
	for i, row := range matrix {
		for j := range row {
			if flashed[i][j] {
				matrix[i][j] = 0
				countFlashed++
			}
		}
	}

	return matrix, flashes, countFlashed == len(matrix)*len(matrix[0])
}
