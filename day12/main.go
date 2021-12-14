package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])
	caveMap := createCaveMap(lines)

	fmt.Println("result 1", countPaths(caveMap, 1))
	fmt.Println("result 2", countPaths(caveMap, 2))

}

type Cave struct {
	label string
}

func (c *Cave) isStart() bool {
	return c.label == "start"
}

func (c *Cave) isEnd() bool {
	return c.label == "end"
}

func (c *Cave) isSmall() bool {
	return c.label == strings.ToLower(c.label)
}

func createCaveMap(lines []string) map[Cave][]Cave {
	res := make(map[Cave][]Cave)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			a := parts[0]
			b := parts[1]
			caves := []Cave{
				{a}, {b},
			}
			caveOne, caveTwo := caves[0], caves[1]
			res[caveOne] = append(res[caveOne], caveTwo)
			res[caveTwo] = append(res[caveTwo], caveOne)
		}
	}

	return res
}

func countPaths(caveMap map[Cave][]Cave, limit int) int {
	s := Cave{"start"}
	visited := map[Cave]int{}
	paths := 0

	traverse(s, caveMap, limit, &paths, &visited)
	return paths
}

func traverse(s Cave, cm map[Cave][]Cave, smallCaveMaxVisit int, paths *int, visited *map[Cave]int) {
	if s.isEnd() {
		*paths++
		return
	}

	if s.isSmall() {
		(*visited)[s]++

		visitedSmallCaves := 0
		for cave, _ := range *visited {
			if (*visited)[cave] > 1 {
				visitedSmallCaves++
			}

			if (*visited)[cave] > smallCaveMaxVisit {
				(*visited)[s]--
				return
			}
		}

		if visitedSmallCaves > 1 {
			(*visited)[s]--
			return
		}
	}

	for _, d := range cm[s] {
		if d.isStart() {
			continue
		}

		traverse(d, cm, smallCaveMaxVisit, paths, visited)
	}

	if s.isSmall() {
		(*visited)[s]--
	}
}
