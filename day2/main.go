package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type submarine1 struct {
	depth    int
	position int
}

type submarine2 struct {
	aim      int
	depth    int
	position int
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sub1 := newSubmarine1()
	sub2 := newSubmarine2()
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		cmd := parts[0]
		param, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		switch cmd {
		case "forward":
			sub1.forward(param)
			sub2.forward(param)
		case "up":
			sub1.up(param)
			sub2.up(param)
		case "down":
			sub1.down(param)
			sub2.down(param)
		}
	}
	fmt.Printf("submarine 1 result = %d\n", sub1.depth*sub1.position)
	fmt.Printf("submarine 2 result = %d\n", sub2.depth*sub2.position)
}

func newSubmarine1() *submarine1 {
	return &submarine1{
		depth:    0,
		position: 0,
	}
}

func newSubmarine2() *submarine2 {
	return &submarine2{
		aim:      0,
		depth:    0,
		position: 0,
	}
}

func (sub *submarine1) up(depth int) {
	sub.depth -= depth
}

func (sub *submarine1) down(depth int) {
	sub.depth += depth
}

func (sub *submarine1) forward(x int) {
	sub.position += x
}

func (sub *submarine2) up(x int) {
	sub.aim -= x
}

func (sub *submarine2) down(x int) {
	sub.aim += x
}

func (sub *submarine2) forward(x int) {
	sub.position += x
	sub.depth += sub.aim * x
}
