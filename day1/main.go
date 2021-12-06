package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	values := []int{}
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}
	fmt.Printf("%d\n", SonarSweep(values))
	fmt.Printf("%d\n", SonarWindowSweep(values))
}

func SonarSweep(depths []int) int {
	count := 0
	for i, depth := range depths {
		if i > 0 && depth > depths[i-1] {
			count++
		}
	}
	return count
}

func SonarWindowSweep(depths []int) int {
	sums := []int{}
	for i := 0; i < len(depths)-2; i++ {
		sum := depths[i] + depths[i+1] + depths[i+2]
		sums = append(sums, sum)
	}
	return SonarSweep(sums)
}
