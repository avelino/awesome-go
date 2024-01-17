package main

import "fmt"

type Graph struct {
	graph map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		graph: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(node1, node2 int) {
	g.graph[node1] = append(g.graph[node1], node2)
}

func (g *Graph) AddNode(node int) {
	if _, ok := g.graph[node]; !ok {
		g.graph[node] = []int{}
	}
}

func (g *Graph) DFS(startNode int) {
	visited := make(map[int]bool)
	g.dfsHelper(startNode, visited)
}

func (g *Graph) dfsHelper(node int, visited map[int]bool) {
	visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range g.graph[node] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited)
		}
	}
}

func (g *Graph) BFS(startNode int) {
	visited := make(map[int]bool)
	queue := []int{startNode}
	visited[startNode] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)

		for _, neighbor := range g.graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	// Create a graph
	graph := NewGraph()

	// Add nodes
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)

	// Add edges
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)

	// Perform DFS traversal
	fmt.Println("DFS traversal:")
	graph.DFS(1)
	fmt.Println()

	// Perform BFS traversal
	fmt.Println("BFS traversal:")
	graph.BFS(1)
	fmt.Println()
}