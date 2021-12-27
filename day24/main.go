package main

// derived from https://www.reddit.com/r/adventofcode/comments/rnejv5/2021_day_24_solutions/hpv7g7j/?utm_source=reddit&utm_medium=web2x&context=3
import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

type Stack []Tuple

func (s *Stack) Push(v Tuple) {
	*s = append(*s, v)
}

func (s *Stack) Pop() Tuple {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}

type Tuple struct{ a, b int }

func main() {
	lines := shared.ReadLines(os.Args[1])
	dig := 0
	stack := Stack{}
	digits := map[int]Tuple{}
	var push bool
	var sub int

	for i, line := range lines {
		op := strings.Split(line, " ")
		operands := op[1:]
		if i%18 == 4 {
			push = operands[1] == "1"
		}
		if i%18 == 5 {
			n, e := strconv.Atoi(operands[1])
			if e != nil {
				fmt.Println(i, line)
				log.Fatal(e)
			}
			sub = n
		}
		if i%18 == 15 {
			if push {
				n, e := strconv.Atoi(operands[1])
				if e != nil {
					fmt.Println(i, line)
					log.Fatal(e)
				}
				stack.Push(Tuple{dig, n})
			} else {
				t := stack.Pop()
				sibling, add := t.a, t.b
				diff := add + sub
				if diff < 0 {
					digits[sibling] = Tuple{-diff + 1, 9}
					digits[dig] = Tuple{1, 9 + diff}
				} else {
					digits[sibling] = Tuple{1, 9 - diff}
					digits[dig] = Tuple{1 + diff, 9}
				}
			}
			dig++
		}
	}
	keys := []int{}
	for k := range digits {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	max := ""
	min := ""
	for _, d := range keys {
		min += string(rune('0' + digits[d].a))
		max += string(rune('0' + digits[d].b))

	}
	fmt.Println(max)
	fmt.Println(min)
}
