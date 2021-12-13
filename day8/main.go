package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	inputs, outputs := parseLines(lines)
	result := countUniqueDigits(outputs)
	fmt.Printf("result 1 = %d\n", result)
	displays := makeDisplays(inputs, outputs)
	fmt.Printf("Display = %v\n", displays)
	fmt.Printf("result 2 = %d\n", sum(displays))
}

func parseLines(lines []string) (inputs, outputs []string) {
	inputs = []string{}
	outputs = []string{}
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		if len(parts) == 2 {
			inputs = append(inputs, parts[0])
			outputs = append(outputs, parts[1])
		}
	}
	return inputs, outputs
}

func countUniqueDigits(display []string) int {
	count := 0
	for _, d := range display {
		digits := strings.Split(d, " ")
		for _, digit := range digits {
			switch len(digit) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func makeDisplays(inputs, outputs []string) []int {
	result := []int{}
	for i, input := range inputs {
		display := makeDisplay(input)
		value := calcValue(outputs[i], display)
		fmt.Printf("i = %s, display = %v\n", input, display)
		fmt.Printf("o = %s, value = %d\n", outputs[i], value)
		result = append(result, value)
	}
	return result
}

func makeDisplay(input string) (result map[string]int) {
	result = map[string]int{}
	found := map[int]string{}
	inputs := strings.Split(input, " ")
	for _, d := range inputs {
		k := SortString(d)
		switch len(d) {
		case 2:
			result[k] = 1
			found[1] = k
		case 4:
			result[k] = 4
			found[4] = k
		case 3:
			result[k] = 7
			found[7] = k
		case 7:
			result[k] = 8
			found[8] = k
		}
	}
	inputs = reduceInputs(inputs, result)
	for _, d := range inputs {
		k := SortString(d)
		four, ok4 := found[4]
		one, ok1 := found[1]

		switch len(d) {
		case 5:
			if ok1 && ok4 {
				l1 := lettersIn(k, one)
				l4 := lettersIn(k, four)
				if l1 == 1 && l4 == 3 {
					found[5] = k
					result[k] = 5
				} else if l1 == 2 && l4 == 3 {
					found[3] = k
					result[k] = 3
				} else if l4 == 2 {
					found[2] = k
					result[k] = 2
				}
			}
		case 6:
			if ok1 && ok4 {
				l1 := lettersIn(k, one)
				l4 := lettersIn(k, four)
				if l1 == 1 {
					found[6] = k
					result[k] = 6
				} else if l4 == 3 {
					found[0] = k
					result[k] = 0
				} else if l4 == 4 {
					found[9] = k
					result[k] = 9
				}
			}
		}
	}

	return
}

func reduceInputs(inputs []string, result map[string]int) []string {
	newInputs := []string{}
	for _, i := range inputs {
		if _, ok := result[i]; !ok {
			newInputs = append(newInputs, i)
		}
	}
	return newInputs
}

func lettersIn(str, substr string) int {
	count := 0
	for _, c := range substr {
		if strings.ContainsRune(str, c) {
			count++
		}
	}
	return count
}

func calcValue(output string, display map[string]int) int {
	result := 0
	for _, o := range strings.Split(output, " ") {
		if v, ok := display[SortString(o)]; ok {
			result = result*10 + v
		}
	}
	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func sum(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
