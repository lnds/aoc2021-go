package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

type Point struct {
	x, y, z int
}

type Scanner struct {
	id        int
	points    []Point
	distances [][]int
	offset    Point
	axisSign  []int
	axisMap   Point
}

func newScanner(id int) *Scanner {
	return &Scanner{
		id:        id,
		points:    []Point{},
		distances: [][]int{},
		offset:    Point{0, 0, 0},
		axisSign:  []int{1, 1, 1},
		axisMap:   Point{0, 1, 2},
	}
}

func (s *Scanner) addPoints(points []Point) {
	s.points = append([]Point{}, points...)
}

func (s *Scanner) calculateDistances() {
	s.distances = [][]int{}
	for i := range s.points {
		p1 := s.points[i]
		distances := []int{}

		for j := range s.points {
			if i == j {
				continue
			}
			p2 := s.points[j]
			dist := (p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y) + (p2.z-p1.z)*(p2.z-p1.z)
			distances = append(distances, dist)
		}
		sort.Ints(distances)
		s.distances = append(s.distances, distances)
	}
}

func (s *Scanner) findOverlapping(other *Scanner) bool {
	commonPoints := map[int]int{}
	for i := range s.distances {
		for j := range other.distances {
			c := 0
			if _, ok := commonPoints[i]; ok {
				break
			}

			for k := 0; k < 12; k++ {
				if s.distances[i][k] == other.distances[j][k] {
					c++
				}
			}
			if c > 1 {
				commonPoints[i] = j
			}
		}
	}
	if len(commonPoints) < 1 {
		return false
	}
	fmt.Printf("find overlapping !!\n")

	axisMap := []int{0, 1, 2}
	axisSign := []int{1, 1, 1}
	offset := []int{math.MinInt, math.MinInt, math.MinInt}

	for i := 0; i < 3; i++ {
		for _, j := range []int{1, -1} {
			offx := []int{}
			offy := []int{}
			offz := []int{}

			for key := range commonPoints {
				p1 := s.points[key]
				p2 := other.points[commonPoints[key]]
				pp2 := []int{p2.x, p2.y, p2.z}
				offx = append(offx, p1.x-j*pp2[i])
				offy = append(offy, p1.y-j*pp2[i])
				offz = append(offz, p1.z-j*pp2[i])
			}

			setofx := map[int]int{}
			for _, ox := range offx {
				setofx[ox]++
			}
			if len(setofx) == 1 {
				axisMap[0] = i
				axisSign[0] = j
				offset[0] = offx[0]
			}

			setofy := map[int]int{}
			for _, oy := range offy {
				setofy[oy]++
			}
			if len(setofy) == 1 {
				axisMap[1] = i
				axisSign[1] = j
				offset[1] = offy[0]
			}

			setofz := map[int]int{}
			for _, oz := range offz {
				setofz[oz]++
			}
			if len(setofz) == 1 {
				axisMap[2] = i
				axisSign[2] = j
				offset[2] = offz[0]
			}
		}
	}

	for _, o := range offset {
		if o == math.MinInt {
			return false
		}
	}

	other.offset = Point{offset[0], offset[1], offset[2]}
	other.axisMap = Point{axisMap[0], axisMap[1], axisMap[2]}
	other.axisSign = axisSign
	other.alignPoints()
	return true
}

func (s *Scanner) alignPoints() {
	for i := range s.points {
		x, y, z := s.axisMap.x, s.axisMap.y, s.axisMap.z
		sx, sy, sz := s.axisSign[0], s.axisSign[1], s.axisSign[2]

		spoints := []int{s.points[i].x, s.points[i].y, s.points[i].z}
		newX := s.offset.x + sx*spoints[x]
		newY := s.offset.y + sy*spoints[y]
		newZ := s.offset.z + sz*spoints[z]

		s.points[i].x = newX
		s.points[i].y = newY
		s.points[i].z = newZ
	}
}

const Prefix = "--- scanner "
const Suffix = " ---"

func main() {
	lines := shared.ReadLines(os.Args[1])
	var scanner *Scanner
	points := []Point{}
	scanners := []*Scanner{}
	for _, line := range lines {
		if strings.HasPrefix(line, Prefix) && strings.HasSuffix(line, Suffix) {
			sid := strings.TrimSuffix(strings.TrimPrefix(line, Prefix), Suffix)
			id, _ := strconv.Atoi(sid)
			scanner = newScanner(id)
		} else if line == "" {
			scanner.addPoints(points)
			scanner.calculateDistances()
			scanners = append(scanners, scanner)
			points = []Point{}
		} else {
			p := strings.Split(line, ",")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			z, _ := strconv.Atoi(p[2])
			points = append(points, Point{x, y, z})
		}
	}
	if scanner != nil {
		scanner.addPoints(points)
		scanner.calculateDistances()
		scanners = append(scanners, scanner)
	}

	fmt.Printf("scanners = %v#\n", scanners)

	processed := []*Scanner{scanners[0]}
	toProcess := scanners[1:]
	for len(toProcess) > 0 {
		fmt.Printf("len to process = %d\n", len(toProcess))
		scanner = toProcess[0]
		toProcess = toProcess[1:]

		var ok bool
		for i, aligned := range processed {
			ok = aligned.findOverlapping(scanner)
			if ok {
				fmt.Printf("Aligned to {%d}\n", i)
				break
			}
		}

		if ok {
			processed = append(processed, scanner)
		} else {
			toProcess = append(toProcess, scanner)
		}
	}

	points = []Point{}
	for _, scanner := range processed {
		for _, point := range scanner.points {
			if !in(point, points) {
				points = append(points, point)
			}
		}
	}
	fmt.Printf("beacons = %d\n", len(points))

	maxd := 0
	for i := range processed {
		for j := i; j < len(processed); j++ {
			s1 := processed[i]
			s2 := processed[j]
			dist := abs(s1.offset.x-s2.offset.x) + abs(s1.offset.y-s2.offset.y) + abs(s1.offset.z-s2.offset.z)
			if dist > maxd {
				maxd = dist
			}
		}
	}
	fmt.Printf("max manhattan distance = %d\n", maxd)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func in(point Point, points []Point) bool {
	for _, p := range points {
		if p.x == point.x && p.y == point.y {
			return true
		}
	}
	return false
}
