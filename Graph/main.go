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

func (g *Graph) BFS(start int) {
	seen := make(map[int]bool)
	queue := []int{}
	queue=append(queue, start)
	seen[start] = true
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", vertex)
		for _, neighbour := range g.List[vertex] {
			if !seen[neighbour] {
				seen[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	fmt.Println()
}

func (g *Graph) DFS(start int){
	seen:=make(map[int]bool)
	stack:=[]int{start}
	for len(stack)>0{
		vertex:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		if !seen[vertex]{
			seen[vertex]=true
			fmt.Printf("%d ",vertex)
			for _, neighbours:= range g.List[vertex]{
				if !seen[neighbours]{
					stack=append(stack, neighbours)
				}
			}
		}
	}
	fmt.Println()
}

func main() {
	graph := &Graph{}
	
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)


	graph.AddEdge(1, 5)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 1)
	graph.Print()


	fmt.Print("BFS: ")
	graph.BFS(1)
	fmt.Print("DFS: ")
	graph.DFS(1)
}
