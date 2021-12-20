package main

import (
	"fmt"
	"os"

	"github.com/lnds/aoc2021-go/shared"
)

type Value struct {
	value int
	depth int
}

type Number []Value

func (n Number) Magnitude() int {
	curr := append(Number{}, n...)
	last := Number{}
	for !equal(curr, last) {
		deepest := 0
		for _, v := range curr {
			if v.depth > deepest {
				deepest = v.depth
			}
		}

		last = append(Number{}, curr...)

		if len(curr) == 1 {
			break
		}

		for i := 0; i < len(curr)-1; i++ {
			if curr[i].depth == deepest {
				left := 3 * curr[i].value
				right := 2 * curr[i+1].value
				curr[i].value = left + right
				curr[i].depth--
				p0 := append(Number{}, curr[0:i+1]...)
				p1 := append(Number{}, curr[i+2:]...)
				curr = append(Number{}, p0...)
				curr = append(curr, p1...)
				break
			}
		}
	}

	return curr[0].value
}

func addition(numbers []Number) Number {
	acum := append(Number{}, numbers[0]...)
	for _, num := range numbers[1:] {
		m := add(acum, num)
		l := Number{}
		for !equal(m, l) {
			reduced := reduce(m)
			l = append(Number{}, m...)
			m = append(Number{}, reduced...)
		}
		acum = append(Number{}, m...)
	}
	return acum
}

func reduce(num Number) Number {
	result := append(Number{}, num...)
	for i := 0; i < len(num); i++ {
		value := num[i].value
		depth := num[i].depth

		if depth > 4 {
			// explode
			if i > 0 {
				result[i-1].value += value
			}
			if i+2 < len(num) {
				result[i+2].value += result[i+1].value
			}
			p0 := append(Number{}, result[0:i]...)
			p1 := append(Number{}, result[i+1:]...)
			result = append(Number{}, p0...)
			result = append(result, p1...)
			if i < len(result) {
				result[i].value = 0
				result[i].depth = depth - 1
			}

			return result
		}
	}

	for i := 0; i < len(num); i++ {
		value := num[i].value
		depth := num[i].depth
		if value >= 10 {
			vhd := value / 2
			vhu := value/2 + value%2
			result[i].value = vhd
			result[i].depth++
			p0 := append(Number{}, result[0:i+1]...)
			p1 := append(Number{}, result[i+1:]...)
			result = append(Number{}, p0...)
			result = append(result, Value{vhu, depth + 1})
			result = append(result, p1...)
			return result
		}
	}
	return result
}

func equal(a, b Number) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func add(n1, n2 Number) Number {
	result := Number{}
	for _, v := range n1 {
		result = append(result, Value{v.value, v.depth + 1})
	}
	for _, v := range n2 {
		result = append(result, Value{v.value, v.depth + 1})
	}
	return result
}

func main() {
	lines := shared.ReadLines(os.Args[1])
	numbers := parseNumbers(lines)
	for _, n := range numbers {
		fmt.Printf("|%v| = %d\n", n, n.Magnitude())
	}
	s := addition(numbers)
	fmt.Printf("addition: |%v| = %d\n", s, s.Magnitude())

	max := 0
	for _, a := range numbers {
		for _, b := range numbers {
			nums := []Number{a, b}
			s := addition(nums)
			m := s.Magnitude()
			if m > max {
				max = m
			}
		}
	}
	fmt.Println("result 2 = ", max)
}

func parseNumbers(lines []string) []Number {
	result := []Number{}
	for _, line := range lines {
		if line != "" {
			result = append(result, parseNumber(line))

		}
	}
	return result
}

func parseNumber(line string) Number {
	depth := 0
	i := 0
	result := Number{}
	for i < len(line) {
		c := line[i]
		switch c {
		case '[':
			depth++
			i++
		case ']':
			depth--
			i++
		case ',':
			i++
		default:
			v0 := int(c - '0')
			v1 := int(line[i+1] - '0')
			var v int
			if v1 >= 0 && v1 <= 9 {
				v = v0*10 + v1
				i++
			} else {
				v = v0
			}
			i++
			result = append(result, Value{v, depth})
		}
	}
	return result
}
