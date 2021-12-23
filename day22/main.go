package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

type Pos struct{ x, y, z int }

type Cuboid struct {
	lo   Pos
	hi   Pos
	sign int
}

func newCuboid(lo, hi Pos, count int) Cuboid {
	return Cuboid{lo, hi, count}
}

func (c Cuboid) volume() int {
	return (c.hi.x - c.lo.x + 1) * (c.hi.y - c.lo.y + 1) * (c.hi.z - c.lo.z + 1)
}

func intersect(c1, c2 Cuboid) bool {
	if c1.lo.x > c2.hi.x || c1.hi.x < c2.lo.x {
		return false
	}
	if c1.lo.y > c2.hi.y || c1.hi.y < c2.lo.y {
		return false
	}
	if c1.lo.z > c2.hi.z || c1.hi.z < c2.lo.z {
		return false
	}
	return true
}

func calcIntersection(c1, c2 Cuboid) Cuboid {
	minX := max(c1.lo.x, c2.lo.x)
	maxX := min(c1.hi.x, c2.hi.x)

	minY := max(c1.lo.y, c2.lo.y)
	maxY := min(c1.hi.y, c2.hi.y)

	minZ := max(c1.lo.z, c2.lo.z)
	maxZ := min(c1.hi.z, c2.hi.z)

	sign := c1.sign * c2.sign
	if c1.sign == c2.sign {
		sign = -c1.sign
	} else if c1.sign == 1 && c2.sign == -1 {
		sign = 1
	}
	return newCuboid(Pos{minX, minY, minZ}, Pos{maxX, maxY, maxZ}, sign)
}

func main() {
	lines := shared.ReadLines(os.Args[1])

	result := processLines1(lines, true)
	fmt.Println("solution 1", result)

	result = processLines2(lines, true)
	fmt.Println("solution 2", result)
}

func processLines1(lines []string, limit bool) int {

	grid := map[Pos]bool{}

	for _, step := range lines {
		on := step[0:2] == "on"
		pos := strings.Index(step, " ")
		p := strings.Split(step[pos+1:], ",")
		xs := strings.Split(strings.TrimPrefix(p[0], "x="), "..")
		xinf, _ := strconv.Atoi(xs[0])
		xsup, _ := strconv.Atoi(xs[1])
		ys := strings.Split(strings.TrimPrefix(p[1], "y="), "..")
		yinf, _ := strconv.Atoi(ys[0])
		ysup, _ := strconv.Atoi(ys[1])
		zs := strings.Split(strings.TrimPrefix(p[2], "z="), "..")
		zinf, _ := strconv.Atoi(zs[0])
		zsup, _ := strconv.Atoi(zs[1])

		if xinf < -50 || xsup > 50 || yinf < -50 || ysup > 50 || zinf < -50 || zsup > 50 {
			continue
		}
		for x := xinf; x <= xsup; x++ {
			for y := yinf; y <= ysup; y++ {
				for z := zinf; z <= zsup; z++ {
					if on {
						grid[Pos{x, y, z}] = true
					} else {
						delete(grid, Pos{x, y, z})
					}
				}
			}
		}
	}
	return len(grid)
}

func processLines2(lines []string, limit bool) int {

	cuboids := []Cuboid{}

	for _, step := range lines {
		on := step[0:2] == "on"
		pos := strings.Index(step, " ")
		p := strings.Split(step[pos+1:], ",")
		xs := strings.Split(strings.TrimPrefix(p[0], "x="), "..")
		xinf, _ := strconv.Atoi(xs[0])
		xsup, _ := strconv.Atoi(xs[1])
		ys := strings.Split(strings.TrimPrefix(p[1], "y="), "..")
		yinf, _ := strconv.Atoi(ys[0])
		ysup, _ := strconv.Atoi(ys[1])
		zs := strings.Split(strings.TrimPrefix(p[2], "z="), "..")
		zinf, _ := strconv.Atoi(zs[0])
		zsup, _ := strconv.Atoi(zs[1])

		sign := -1
		if on {
			sign = 1
		}
		curr := Cuboid{Pos{xinf, yinf, zinf}, Pos{xsup, ysup, zsup}, sign}
		intersections := []Cuboid{}

		for _, cuboid := range cuboids {
			if intersect(curr, cuboid) {
				intersection := calcIntersection(curr, cuboid)
				intersections = append(intersections, intersection)
			}
		}

		for _, i := range intersections {
			cuboids = append(cuboids, i)
		}
		if on {
			cuboids = append(cuboids, curr)
		}
	}
	res := 0
	for _, cuboid := range cuboids {
		res += cuboid.volume() * cuboid.sign
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
