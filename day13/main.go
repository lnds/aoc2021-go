package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])
	coordinates, folds := parseLines(lines)

	paper := buildPaper(coordinates)

	paper = fold(paper, folds, 1)

	visibleDots := countVisibleDots(paper)

	fmt.Printf("visibleDots = %d\n", visibleDots)

	paper = fold(paper, folds[1:], len(folds)-1)

	fmt.Printf("paper = %v\n", paper)

	printPaper(paper)
}

type Point struct {
	x, y int
}

const (
	AlongX = iota
	AlongY
)

type Fold struct {
	axis int
	pos  int
}

func parseLines(lines []string) ([]Point, []Fold) {
	points := []Point{}
	last := 0
	for l, line := range lines {
		last = l
		point, ok := parsePoint(line)
		if !ok {
			break
		} else {
			points = append(points, point)
		}
	}
	folds := []Fold{}
	for _, line := range lines[last+1:] {
		fold, ok := parseFold(line)
		if !ok {
			break
		} else {
			folds = append(folds, fold)
		}
	}
	return points, folds
}

func parsePoint(line string) (Point, bool) {
	if line == "" {
		return Point{}, false
	}
	p := strings.Split(line, ",")
	if len(p) != 2 {
		return Point{}, false
	}
	p0, _ := strconv.Atoi(p[0])
	p1, _ := strconv.Atoi(p[1])
	return Point{p0, p1}, true
}

func parseFold(line string) (Fold, bool) {
	line = strings.TrimPrefix(line, "fold along ")
	p := strings.Split(line, "=")
	if len(p) != 2 {
		return Fold{}, false
	}
	value, _ := strconv.Atoi(p[1])
	switch p[0] {
	case "x":
		return Fold{AlongX, value}, true
	case "y":
		return Fold{AlongY, value}, true
	default:
		return Fold{}, false
	}
}

type Paper [][]bool

func buildPaper(coords []Point) Paper {
	var maxX, maxY int
	for _, c := range coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	paper := make(Paper, maxY+1)
	for y := range paper {
		paper[y] = make([]bool, maxX+1)
	}

	for _, c := range coords {
		paper[c.y][c.x] = true
	}

	return paper
}

func fold(paper Paper, folds []Fold, limit int) Paper {
	for i := 0; i < limit; i++ {
		switch folds[i].axis {
		case AlongX:
			paper = foldX(paper, folds[i].pos)
		case AlongY:
			paper = foldY(paper, folds[i].pos)
		}
	}
	return paper
}

func foldY(paper Paper, pos int) Paper {
	newPaper := make(Paper, pos)
	for y := range newPaper {
		newPaper[y] = make([]bool, len(paper[0]))
	}

	newY := pos - 1
	for y := pos + 1; y < len(paper); y++ {
		for x := 0; x < len(paper[y]); x++ {
			newPaper[newY][x] = paper[newY][x] || paper[y][x]
		}
		newY--
	}
	return newPaper
}

func foldX(paper Paper, pos int) Paper {
	newPaper := make(Paper, len(paper))
	for y := range newPaper {
		newPaper[y] = make([]bool, pos)
	}
	for y := range paper {
		newX := pos - 1
		for x := pos + 1; newX >= 0 && x < len(paper[y]); x++ {
			newPaper[y][newX] = paper[y][newX] || paper[y][x]
			newX--
		}
	}
	return newPaper
}

func countVisibleDots(paper Paper) (count int) {
	for y := range paper {
		for x := range paper[y] {
			if paper[y][x] {
				count++
			}
		}
	}
	return
}

func printPaper(paper Paper) {
	for y := range paper {
		for x := range paper[y] {
			if paper[y][x] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
