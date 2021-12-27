package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"

	"github.com/lnds/aoc2021-go/shared"
)

type Amphipod = rune

const (
	Amber     = 'A'
	Bronze    = 'B'
	Copper    = 'C'
	Desert    = 'D'
	Forbidden = '#'
	Empty     = '.'
)

type Square Amphipod

type State struct {
	cost    int
	hallway []Square
	halls   [][]Amphipod
}

func clone(s State) (result State) {
	result.cost = s.cost
	result.hallway = []Square{}
	result.hallway = append(result.hallway, s.hallway...)
	result.halls = make([][]Amphipod, len(s.halls))
	for j, h := range s.halls {
		result.halls[j] = append([]Amphipod{}, h...)
	}
	return
}

func parseState(lines []string) (State, int) {
	hallway := []Square{}
	for _, c := range lines[1] {
		switch c {
		case '#':
			hallway = append(hallway, Forbidden)
		case '.':
			hallway = append(hallway, Empty)

		case 'A':
			hallway = append(hallway, Amber)
		case 'B':
			hallway = append(hallway, Bronze)
		case 'C':
			hallway = append(hallway, Copper)
		case 'D':
			hallway = append(hallway, Desert)
		default:
			panic("wtf")
		}
	}
	hallway[3] = Forbidden
	hallway[5] = Forbidden
	hallway[7] = Forbidden
	hallway[9] = Forbidden

	halls := make([][]Amphipod, 4)
	capacity := 0
outer:
	for _, line := range lines[2:] {
		for i := 0; i < 4; i++ {
			pos := i*2 + 3
			switch line[pos] {
			case '#':
				break outer
			case 'A':
				halls[i] = append([]Amphipod{Amber}, halls[i]...)
			case 'B':
				halls[i] = append([]Amphipod{Bronze}, halls[i]...)
			case 'C':
				halls[i] = append([]Amphipod{Copper}, halls[i]...)
			case 'D':
				halls[i] = append([]Amphipod{Desert}, halls[i]...)
			default:
				panic("wtf")
			}
		}
		capacity++
	}
	return State{cost: 0, hallway: hallway, halls: halls}, capacity
}

func (s State) Key() string {
	result := "("
	for _, c := range s.hallway {
		result += string(c)
	}
	result += ",["
	for _, h := range s.halls {
		result += "[" + string(h) + "],"
	}
	result += "])"
	return result
}

func (s State) String() string {
	lines := []string{}

	hs := ""
	for _, c := range s.hallway {
		hs += string(c)
	}
	lines = append(lines, hs)
	max := 0
	for _, h := range s.halls {
		if len(h) > max {
			max = len(h)
		}
	}
	for i := 0; i < max; i++ {
		lines = append(lines, "  #")
	}
	for j := 0; j < 4; j++ {
		for i := 1; i < max+1; i++ {
			if i-1 < len(s.halls[j]) {
				lines[i] += string(s.halls[j][i-1]) + "#"
			} else {
				lines[i] += ".#"
			}
		}
	}
	return strings.Join(lines, "\n")
}

func isFinal(state State, depth int) bool {
	return len(filterEqual(state.halls[0], Amber)) == depth &&
		len(filterEqual(state.halls[1], Bronze)) == depth &&
		len(filterEqual(state.halls[2], Copper)) == depth &&
		len(filterEqual(state.halls[3], Desert)) == depth
}

var PosJ = map[Amphipod]int{
	Amber:  0,
	Bronze: 1,
	Copper: 2,
	Desert: 3,
}

var Costs = map[Amphipod]int{
	Amber:  1,
	Bronze: 10,
	Copper: 100,
	Desert: 1000,
}

type KeySet map[string]struct{}

func (s KeySet) insert(key string) {
	s[key] = struct{}{}
}

func (s KeySet) contains(key string) bool {
	_, ok := s[key]
	return ok
}

func main() {
	lines := shared.ReadLines(os.Args[1])
	source, capacity := parseState(lines)

	seen := KeySet{}

	binReqs := []Amphipod{Amber, Bronze, Copper, Desert}

	q := newPriorityQueue()
	heap.Push(q, source)
	for q.Len() > 0 {
		state := heap.Pop(q).(State)
		index := state.Key()

		if seen.contains(index) {
			continue
		}
		if isFinal(state, capacity) {
			fmt.Println("final, cost = ", state.cost)
			break
		}
		seen.insert(index)

		added := 0
		for i, n := range state.hallway {
			if amphipod, ok := ocupied(n); ok {
				targetHallIndex := PosJ[amphipod]
				targetHall := state.halls[targetHallIndex]
				if len(targetHall) == 0 || len(filterDistinct(targetHall, amphipod)) == 0 {
					reachable := true
					ti := (targetHallIndex+1)*2 + 1

					if i < ti {
						for k := i + 1; k < ti+1; k++ {
							if state.hallway[k] != Forbidden && state.hallway[k] != Empty {
								reachable = false
							}
						}
					} else {
						for k := i - 1; k >= ti; k-- {
							if state.hallway[k] != Forbidden && state.hallway[k] != Empty {
								reachable = false
							}
						}
					}
					if reachable {
						steps := abs((targetHallIndex+1)*2+1-i) + (capacity - len(targetHall))
						score := Costs[amphipod] * steps
						c := Costs[amphipod]
						if c != 1 && c != 10 && c != 100 && c != 1000 {
							fmt.Println("WTF", c)
						}
						if steps <= 0 || steps >= 12 {
							fmt.Println("WTF STEPS", steps)
						}
						newState := clone(state)
						newState.hallway[i] = Empty
						newState.halls[targetHallIndex] = append(newState.halls[targetHallIndex], amphipod)
						newState.cost += score
						heap.Push(q, newState)

						added++
					}
				}
			}
		}

		for i, hall := range state.halls {
			if len(hall) > 0 {
				topItem := hall[len(hall)-1]
				if topItem != binReqs[i] || len(filterDistinct(hall, binReqs[i])) > 0 {
					cp := (i+1)*2 + 1
					destinations := []int{}
					for n := cp - 1; n >= 1; n-- {
						if state.hallway[n] == Empty {
							destinations = append(destinations, n)
						} else if state.hallway[n] != Forbidden {
							break
						}
					}
					for n := cp; n < 12; n++ {
						if state.hallway[n] == Empty {
							destinations = append(destinations, n)
						} else if state.hallway[n] != Forbidden {
							break
						}
					}
					for _, d := range destinations {
						newState := clone(state)
						cost := capacity - len(hall) + 1
						l := len(newState.halls[i])
						moved := newState.halls[i][l-1]
						newState.halls[i] = newState.halls[i][0 : l-1]
						distance := abs(cp - d)
						cost += distance
						cost = cost * Costs[moved]
						newState.hallway[d] = Square(moved)
						newState.cost += cost
						heap.Push(q, newState)
						added++
					}
				}
			}
		}

	}

}

func ocupied(s Square) (Amphipod, bool) {
	if s == Amber || s == Bronze || s == Copper || s == Desert {
		return Amphipod(s), true
	}
	return Amphipod(Empty), false
}

func filterDistinct(hall []Amphipod, amphipod Amphipod) []Amphipod {
	result := []Amphipod{}
	for i := 0; i < len(hall); i++ {
		if hall[i] != amphipod {
			result = append(result, hall[i])
		}
	}
	return result
}

func filterEqual(hall []Amphipod, amphipod Amphipod) []Amphipod {
	result := []Amphipod{}
	for i := 0; i < len(hall); i++ {
		if hall[i] == amphipod {
			result = append(result, hall[i])
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type PriorityQueue []State

func newPriorityQueue() *PriorityQueue {
	result := &PriorityQueue{}
	heap.Init(result)
	return result
}

func (pq *PriorityQueue) Push(x interface{}) {
	s := x.(State)
	*pq = append(*pq, s)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	s := old[n-1]
	*pq = old[0 : n-1]
	return s
}

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].cost < (*pq)[j].cost
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
