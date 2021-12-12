package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		bits := scanner.Text()
		lines = append(lines, bits)
	}
	count0, count1 := calcFreq01(lines)

	gamma := ""
	epsilon := ""
	for pos := 0; pos < len(count0); pos++ {
		if count0[pos] > count1[pos] {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}

	g := fromBinary(gamma)
	e := fromBinary(epsilon)
	fmt.Printf("gamma   = %s (%d)\n", gamma, g)
	fmt.Printf("epsilon = %s (%d)\n", epsilon, e)
	fmt.Printf("result = %d\n", e*g)

	o2 := calcOxygenGeneratorRating(lines, count0, count1)
	co2 := calcCO2GeneratorRating(lines, count0, count1)
	fmt.Printf("result2 = %d\n", o2*co2)
}

func calcFreq01(lines []string) (count0, count1 map[int]int) {
	count0 = map[int]int{}
	count1 = map[int]int{}
	for _, bits := range lines {
		for pos, bit := range bits {
			switch bit {
			case '0':
				count0[pos]++
			case '1':
				count1[pos]++
			}
		}
	}
	return
}

func calcOxygenGeneratorRating(lines []string, count0, count1 map[int]int) int {
	n := len(count0)
	for pos := 0; pos < n; pos++ {
		if len(lines) == 1 {
			break
		}
		if count0[pos] > count1[pos] {
			newLines := []string{}
			for _, l := range lines {
				if l[pos] == '0' {
					newLines = append(newLines, l)
				}
			}
			lines = newLines
		} else {
			newLines := []string{}
			for _, l := range lines {
				if l[pos] == '1' {
					newLines = append(newLines, l)
				}
			}
			lines = newLines
		}
		count0, count1 = calcFreq01(lines)
	}
	result := fromBinary(lines[0])
	fmt.Printf("oxygen generator rate = %d\n", result)
	return result
}

func calcCO2GeneratorRating(lines []string, count0, count1 map[int]int) int {
	n := len(count0)
	for pos := 0; pos < n; pos++ {
		if len(lines) == 1 {
			break
		}
		if count0[pos] <= count1[pos] {
			newLines := []string{}
			for _, l := range lines {
				if l[pos] == '0' {
					newLines = append(newLines, l)
				}
			}
			lines = newLines
		} else {
			newLines := []string{}
			for _, l := range lines {
				if l[pos] == '1' {
					newLines = append(newLines, l)
				}
			}
			lines = newLines
		}
		count0, count1 = calcFreq01(lines)
	}
	result := fromBinary(lines[0])
	fmt.Printf("co2 generator rate = %d\n", result)
	return result
}

func fromBinary(binary string) int {
	result := 0
	for _, c := range binary {
		if c == '1' {
			result = result*2 + 1
		} else {
			result *= 2
		}
	}
	return result
}
