package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	source      int
	destination int
	weight      int
}

type Graph struct {
	vertices []int
	edges    []Edge
}

func (g *Graph) AddEdge(u, v, w int) {
	edge := Edge{source: u, destination: v, weight: w}
	g.edges = append(g.edges, edge)
}

func (g *Graph) Find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = g.Find(parent, parent[i])
	}
	return parent[i]
}

func (g *Graph) Union(parent []int, rank []int, x, y int) {
	xroot := g.Find(parent, x)
	yroot := g.Find(parent, y)

	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}

func KruskalMST(graph *Graph) {
	vertices := graph.vertices
	edges := graph.edges

	result := make([]Edge, 0)
	parent := make([]int, len(vertices))
	rank := make([]int, len(vertices))

	// Initialize parent and rank
	for i := 0; i < len(vertices); i++ {
		parent[i] = i
		rank[i] = 0
	}

	// Sort edges in ascending order of their weights
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	for _, edge := range edges {
		x := graph.Find(parent, edge.source)
		y := graph.Find(parent, edge.destination)

		if x != y {
			result = append(result, edge)
			graph.Union(parent, rank, x, y)
		}
	}

	// Print the constructed MST
	fmt.Println("Edge \tWeight")
	for _, edge := range result {
		fmt.Printf("%d - %d \t%d\n", edge.source, edge.destination, edge.weight)
	}
}

func main() {
	graph := Graph{
		vertices: []int{0, 1, 2, 3, 4},
	}
	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 3, 6)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(1, 4, 5)
	graph.AddEdge(2, 4, 7)
	graph.AddEdge(3, 4, 9)

	KruskalMST(&graph)
}