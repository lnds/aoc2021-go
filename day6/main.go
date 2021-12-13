package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	fish := parseLine(scanner.Text())

	fmt.Printf("result1 %d\n", countFish(fish, 80))
	fmt.Printf("result2 %d\n", countFish(fish, 256))

}

func countFish(fish []int, days int) int {
	offsprings := make([]int, 9)
	for _, f := range fish {
		offsprings[f]++
	}
	for day := 0; day < days; day++ {
		first := offsprings[0]
		copy(offsprings[0:], offsprings[1:])
		offsprings[8] = first
		offsprings[6] += offsprings[8]
	}
	sum := 0
	for _, f := range offsprings {
		sum += f
	}
	return sum
}
