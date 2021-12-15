package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])
	template, rules := parseRules(lines)

	fmt.Println("result 1 = ", polymerValue(template, rules, 10))

	fmt.Println("result 2 = ", polymerValue(template, rules, 40))

}

func polymerValue(template string, rules Rules, steps int) int {
	var countPairs = make(map[string]int)
	for pos := 0; pos < len(template)-1; pos++ {
		countPairs[template[pos:pos+2]]++
	}
	last := template[len(template)-1]

	for i := 0; i < steps; i++ {
		m := map[string]int{}
		for k, v := range countPairs {
			insert := rules[k]
			first := k[0:1] + insert
			second := insert + k[1:2]

			m[first] += v
			m[second] += v
		}
		countPairs = m
	}

	freq := map[byte]int{}
	freq[last] = 1
	for k, v := range countPairs {
		freq[k[0]] += v
	}
	largest := 0
	smallest := math.MaxInt
	for _, v := range freq {
		norm := v
		if norm > largest {
			largest = norm
		}
		if norm < smallest {
			smallest = norm
		}
	}

	return largest - smallest
}

type Rules map[string]string

func parseRules(lines []string) (string, Rules) {
	template := lines[0]
	rules := Rules{}
	for _, line := range lines[2:] {
		a := strings.Split(line, " -> ")
		rules[a[0]] = a[1]
	}
	return template, rules
}
