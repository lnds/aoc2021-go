package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"

	"github.com/lnds/aoc2021-go/shared"
)

func main() {
	lines := shared.ReadLines(os.Args[1])

	cavern := newMatrix(shared.ParseMatrix09(lines), 1)

	path := minimalPath(cavern)
	fmt.Printf("result 1 = %d\n", path)

	cavern.factor = 5
	path = minimalPath(cavern)
	fmt.Printf("result 2 = %d\n", path)

}

type Matrix struct {
	factor int
	data   shared.Matrix09
}

func newMatrix(data shared.Matrix09, factor int) Matrix {
	return Matrix{factor, data}
}

func (m *Matrix) height() int {
	return len(m.data) * m.factor
}

func (m *Matrix) width() int {
	return len(m.data[0]) * m.factor
}

func (m *Matrix) get(i, j int) int {
	di := i / len(m.data)
	dj := j / len(m.data[0])
	ii := i % len(m.data)
	jj := j % len(m.data[0])
	value := m.data[ii][jj] + di + dj
	if value > 9 {
		return (value - 9) % 10
	} else {
		return value
	}
}

func minimalPath(cavern Matrix) int {
	graph := newGraph()

	fmt.Println("start")
	h := cavern.height()
	w := cavern.width()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			graph.addVertex(Point{i, j})
		}
	}
	fmt.Println("vertex")

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i+1 < h {
				graph.AddEdge(Point{i, j}, Point{i + 1, j})
			}
			if j+1 < w {
				graph.AddEdge(Point{i, j}, Point{i, j + 1})
			}
		}
	}
	fmt.Println("edges")

	return dijkstra(&graph, Point{0, 0}, Point{h - 1, w - 1}, cavern)
}

func dijkstra(graph *Graph, source, target Point, cavern Matrix) int {

	dist := map[Point]int{}
	for v := range graph.adjacency {
		dist[v] = math.MaxInt
	}
	dist[source] = 0

	Q := newPriorityQueue()
	Q.addWithPriority(source, dist[source])

	for !Q.empty() {
		u := Q.extractMin()
		neighbors := graph.adjacency[u]
		for _, v := range neighbors {
			alt := dist[u] + length(u, v, cavern)
			if alt < dist[v] {
				dist[v] = alt
				Q.addWithPriority(v, alt)
			}
		}
	}
	return dist[target]
}

type Item struct {
	point    Point
	priority int
}

type PriorityQueue struct {
	items []Item
}

func newPriorityQueue() *PriorityQueue {
	result := &PriorityQueue{
		items: make([]Item, 0),
	}
	heap.Init(result)
	return result
}

func (pq *PriorityQueue) addWithPriority(p Point, priority int) {
	for _, itm := range pq.items {
		if itm.point == p {
			return
		}
	}
	item := Item{p, priority}
	heap.Push(pq, item)
}

func (pq *PriorityQueue) empty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue) extractMin() Point {
	item := heap.Pop(pq).(Item)
	return item.point
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].priority < pq.items[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]

}

func length(u, v Point, cavern Matrix) int {
	return cavern.get(v.i, v.j)
}

type Point struct {
	i, j int
}

type Graph struct {
	adjacency map[Point][]Point
}

func newGraph() Graph {
	return Graph{
		adjacency: make(map[Point][]Point),
	}
}

func (g *Graph) addVertex(vertex Point) bool {
	if _, ok := g.adjacency[vertex]; ok {
		return false
	}
	g.adjacency[vertex] = []Point{}
	return true
}

func (g *Graph) AddEdge(vertex, node Point) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		return false
	}
	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

func contains(slice []Point, item Point) bool {
	set := make(map[Point]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
