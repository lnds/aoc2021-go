package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseLine(text string) []int {
	parts := strings.Split(text, ",")
	result := []int{}
	for _, n := range parts {
		num, _ := strconv.Atoi(n)
		result = append(result, num)
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
	crabs := parseLine(scanner.Text())
	sort.Ints(crabs)
	fmt.Printf("crabs = %v\n", crabs)
	path, fuel := findPath(crabs)
	fmt.Printf("1. path = %v, fuel = %d\n", path, fuel)

	path, fuel = findPathWithIncreasingCost(crabs)
	fmt.Printf("2. path = %v, fuel = %d\n", path, fuel)
}

func findPath(positions []int) ([]int, int) {
	result := make([]int, len(positions))
	min := sum(positions)
	for v := positions[0]; v <= positions[len(positions)-1]; v++ {
		path := make([]int, len(positions))
		for i, p := range positions {
			d := abs(v - p)
			path[i] = d

		}
		s := sum(path)
		if s < min {
			copy(result[:], path[:])
			min = s
		}
	}
	return result, min
}

func findPathWithIncreasingCost(positions []int) ([]int, int) {
	result := make([]int, len(positions))
	min := math.MaxInt64
	for v := positions[0]; v <= positions[len(positions)-1]; v++ {

		path := make([]int, len(positions))
		for i, p := range positions {
			d := abs(v - p)
			path[i] = (d * (d + 1)) / 2
		}
		s := sum(path)
		if s < min {
			copy(result[:], path[:])
			min = s
		}
	}
	return result, min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
