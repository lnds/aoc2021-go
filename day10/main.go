package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])

	scoreCorrupted := 0
	scoreIncomplete := []int{}
	for _, line := range lines {
		score, stack := parseLine(line)
		scoreCorrupted += score
		if score == 0 && stack != nil {
			score = completeLine(stack)
			scoreIncomplete = append(scoreIncomplete, score)
		}
	}

	sort.Ints(scoreIncomplete)
	m := len(scoreIncomplete) / 2
	fmt.Printf("score corrupted lines = %d\n", scoreCorrupted)
	fmt.Printf("score incomplete lines = %v\n", scoreIncomplete[m])
}

type Stack struct {
	s []rune
}

func newStack() *Stack {
	return &Stack{make([]rune, 0)}
}

func (s *Stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (rune, bool) {
	l := len(s.s)
	if l == 0 {
		return 0, false
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, true
}

func parseLine(line string) (int, *Stack) {
	stack := newStack()
	for _, c := range line {
		switch c {
		case '(', '{', '[', '<':
			stack.Push(c)
		case ')':
			if v, ok := stack.Pop(); ok && v != '(' {
				return 3, nil
			}
		case ']':
			if v, ok := stack.Pop(); ok && v != '[' {
				return 57, nil
			}
		case '}':
			if v, ok := stack.Pop(); ok && v != '{' {
				return 1197, nil
			}
		case '>':
			if v, ok := stack.Pop(); ok && v != '<' {
				return 25137, nil
			}
		}
	}
	return 0, stack
}

func completeLine(stack *Stack) (score int) {
	for c, ok := stack.Pop(); ok; c, ok = stack.Pop() {
		score *= 5
		switch c {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return score
}
