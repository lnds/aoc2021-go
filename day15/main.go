package main

import (
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
	Q := map[Point]bool{}

	dist := map[Point]int{}
	prev := map[Point]*Point{}
	for v, _ := range graph.adjacency {
		dist[v] = math.MaxInt
		prev[v] = nil
		Q[v] = true
	}
	dist[source] = 0

	for len(Q) > 0 {
		du := math.MaxInt
		var u Point
		for v := range Q {
			if dist[v] < du {
				du = dist[v]
				u = v
			}
		}
		delete(Q, u)
		if u == target {
			break
		}

		neighbors := graph.adjacency[u]
		for _, v := range neighbors {
			if _, ok := Q[v]; ok {
				alt := dist[u] + length(u, v, cavern)
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = &u
				}
			}
		}
	}
	return dist[target]
}

func length(u, v Point, cavern Matrix) int {
	if u.i+1 == v.i || u.j+1 == v.j {
		return cavern.get(v.i, v.j)
	}
	return 0
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
