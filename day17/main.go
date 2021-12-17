package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

type Velocity struct {
	vx, vy int
}

type Point struct {
	x, y int
}

type TargetArea [2]Point

func main() {
	lines := shared.ReadLines(os.Args[1])
	targetArea := parseTargetArea(lines[0])

	max := 0
	count := 0
	for vx := 0; vx <= targetArea[1].x; vx++ {
		for vy := targetArea[1].y; vy <= -targetArea[1].y; vy++ {
			v := Velocity{vx, vy}
			h, ok := fly(v, targetArea)
			if ok {
				count++
			}
			if ok && h > max {
				max = h
			}
		}
	}
	fmt.Printf("count = %d\n", count)
	fmt.Printf("highest = %v\n", max)
}

func parseTargetArea(line string) TargetArea {
	line = strings.TrimPrefix(line, "target area: ")
	p := strings.Split(line, ", ")
	px := strings.Split(strings.TrimPrefix(p[0], "x="), "..")
	x0, x1 := shared.ParseInt(px[0]), shared.ParseInt(px[1])
	py := strings.Split(strings.TrimPrefix(p[1], "y="), "..")
	y0, y1 := shared.ParseInt(py[0]), shared.ParseInt(py[1])

	return TargetArea{
		{x0, y1},
		{x1, y0},
	}
}

const (
	Far = iota
	Above
	Inside
	Pass
)

func checkTarget(p Point, target TargetArea) int {
	if p.x < target[0].x && p.y > target[1].y {
		return Far
	}
	if p.x >= target[0].x && p.x <= target[1].x && p.y > target[0].y {
		return Above
	}
	if p.x >= target[0].x && p.x <= target[1].x && p.y <= target[0].y && p.y >= target[1].y {
		return Inside
	}
	return Pass
}

// returns true if reach the target
func fly(v Velocity, target TargetArea) (int, bool) {
	pos := Point{0, 0}
	high := 0
	for {
		if pos.y > high {
			high = pos.y
		}

		switch checkTarget(pos, target) {
		case Inside:
			return high, true
		case Pass:
			return high, false
		default:
			pos.x += v.vx
			pos.y += v.vy
			v.vy--
			if v.vx > 0 {
				v.vx--
			} else if v.vx < 0 {
				v.vx++
			}
		}
	}

}
