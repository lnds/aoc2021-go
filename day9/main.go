package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	matrix := parseMatrix(lines)

	risk, basinRisk := findRiskLevel(matrix)
	fmt.Printf("risk level = %d, basin result= %d\n", risk, basinRisk)
}

type Matrix [][]int

func parseMatrix(lines []string) Matrix {
	matrix := Matrix{}
	for _, line := range lines {
		row := []int{}
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func findRiskLevel(matrix Matrix) (int, int) {
	lowpoints := []int{}
	basins := []int{}
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		for j := 0; j < len(row); j++ {
			adjacents := []int{}
			if i > 0 {
				adjacents = append(adjacents, matrix[i-1][j])
			}
			if j > 0 {
				adjacents = append(adjacents, matrix[i][j-1])
			}
			if i < len(matrix)-1 {
				adjacents = append(adjacents, matrix[i+1][j])
			}
			if j < len(row)-1 {
				adjacents = append(adjacents, matrix[i][j+1])
			}
			c := matrix[i][j]
			n := 0
			for _, v := range adjacents {
				if v > c {
					n++
				}
			}
			if n == len(adjacents) {
				lowpoints = append(lowpoints, c+1)
				basins = append(basins, basinSize(i, j, matrix))
			}
		}
	}

	sort.Ints(basins)
	mult3 := basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]

	return sum(lowpoints), mult3
}

const (
	Undefined = iota
	North
	East
	South
	West
	Halt
)

type Point struct {
	i, j int
}

func basinSize(i, j int, matrix Matrix) (size int) {
	p := Point{i: i, j: j}
	visited := map[Point]bool{p: true}

	q := list.New()
	q.PushBack(p)

	for e := q.Front(); e != nil; e = e.Next() {
		pos := e.Value.(Point)

		height, ok := getHeight(matrix, pos.i, pos.j)
		if !ok || height == 9 {
			continue
		}

		neighbors := []Point{}
		neighbors = append(neighbors, Point{pos.i, pos.j - 1})
		neighbors = append(neighbors, Point{pos.i, pos.j + 1})
		neighbors = append(neighbors, Point{pos.i - 1, pos.j})
		neighbors = append(neighbors, Point{pos.i + 1, pos.j})

		for _, npos := range neighbors {
			if _, ok := visited[npos]; ok {
				continue
			}
			visited[npos] = true
			q.PushBack(npos)
		}

		size++
	}
	return
}

func getHeight(matrix Matrix, i, j int) (int, bool) {
	if i < 0 || i >= len(matrix) {
		return 0, false
	}
	if j < 0 || j >= len(matrix[i]) {
		return 0, false
	}
	return matrix[i][j], true
}

func sum(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
