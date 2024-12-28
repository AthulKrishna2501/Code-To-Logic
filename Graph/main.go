package main

import "fmt"

type Graph struct {
	List map[int][]int
}

func (g *Graph) AddVertex(vertex int) {
	if g.List == nil {
		g.List = make(map[int][]int)
		return
	}

	if _, exists := g.List[vertex]; !exists {
		g.List[vertex] = []int{}
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.List[v1] = append(g.List[v1], v2)
	g.List[v2] = append(g.List[v2], v1)
}

func (g *Graph) Print() {
	for vertex, edge := range g.List {
		fmt.Printf("%d -> %v\n", vertex, edge)
	}
}

func main() {
	graph := &Graph{}
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddEdge(1,5)
	graph.AddEdge(2,4)
	graph.AddEdge(3,1)
	graph.Print()
}
