package main

import (
	"fmt"
	"math"
)

const V = 9

func minDistance(dist []int, sptSet []bool) int {
	min := math.MaxInt32
	var minIndex int

	for v := 0; v < V; v++ {
		if sptSet[v] == false && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

func printSolution(dist []int) {
	fmt.Println("Vertex \t Distance from Source")
	for i := 0; i < V; i++ {
		fmt.Printf("%d \t\t\t\t %d\n", i, dist[i])
	}
}

func dijkstra(graph [V][V]int, src int) {
	dist := make([]int, V)
	sptSet := make([]bool, V)

	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt32
		sptSet[i] = false
	}

	dist[src] = 0

	for count := 0; count < V-1; count++ {
		u := minDistance(dist, sptSet)

		sptSet[u] = true

		for v := 0; v < V; v++ {
			if !sptSet[v] && graph[u][v] != 0 && dist[u] != math.MaxInt32 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	printSolution(dist)
}

func main() {
	graph := [V][V]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}

	src := 0
	dijkstra(graph, src)
}
