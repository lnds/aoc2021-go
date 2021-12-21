package main

import (
	"fmt"
	"os"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])

	algorithm := parseToBinary(lines[0])

	matrix := map[Coord]bool{}

	for r, l := range lines[2:] {
		for c, x := range l {
			if x == '#' {
				matrix[Coord{r, c}] = true
			}
		}
	}

	N := len(lines[2:]) + 1
	m1 := matrix
	m2 := map[Coord]bool{}
	for k, v := range matrix {
		m2[k] = v
	}
	fmt.Println(solve(2, N, m1, algorithm))
	fmt.Println(solve(50, N, m2, algorithm))

}

type Coord struct{ r, c int }

func solve(n int, N int, matrix map[Coord]bool, algorithm []int) int {
	for i := 0; i < n; i++ {
		matrix = step(N, matrix, algorithm)
	}
	count := 0
	for r := -2 - n; r < N+2+n; r++ {
		for c := -2 - n; c < N+2+n; c++ {
			if matrix[Coord{r, c}] {
				count++
			}
		}
	}
	return count
}

func step(N int, matrix map[Coord]bool, algorithm []int) map[Coord]bool {
	set := map[Coord]bool{}
	for r := -N - 1; r < 2*N+1; r++ {
		for c := -N - 1; c < 2*N+1; c++ {
			k := 0
			for _, i := range []int{-1, 0, 1} {
				for _, j := range []int{-1, 0, 1} {
					if _, ok := matrix[Coord{r + i, c + j}]; ok {
						k = k*2 + 1
					} else {
						k *= 2
					}
				}
			}
			if algorithm[k] == 1 {
				set[Coord{r, c}] = true
			}
		}
	}
	return set
}

func parseToBinary(line string) []int {
	result := []int{}
	for _, c := range line {
		if c == '.' {
			result = append(result, 0)
		} else if c == '#' {
			result = append(result, 1)
		}
	}
	return result
}
