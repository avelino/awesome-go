package main

import (
	"fmt"
)

type Edge struct {
	source      int
	destination int
	weight      int
}

func PrimMST(graph [][]int) {
	vertexCount := len(graph)
	parent := make([]int, vertexCount)
	key := make([]int, vertexCount)
	mstSet := make([]bool, vertexCount)

	// Initialize key values as infinity
	for i := 0; i < vertexCount; i++ {
		key[i] = int(^uint(0) >> 1) // set as maximum int value
		mstSet[i] = false
	}

	// Always include the first vertex in MST
	key[0] = 0 // Make key 0 so that this vertex is picked as first vertex
	parent[0] = -1 // First node is always the root of MST

	for count := 0; count < vertexCount-1; count++ {
		u := minKey(key, mstSet, vertexCount)
		mstSet[u] = true

		for v := 0; v < vertexCount; v++ {
			if graph[u][v] != 0 && mstSet[v] == false && graph[u][v] < key[v] {
				parent[v] = u
				key[v] = graph[u][v]
			}
		}
	}

	// Print the constructed MST
	fmt.Println("Edge \tWeight")
	for i := 1; i < vertexCount; i++ {
		fmt.Printf("%d - %d \t%d\n", parent[i], i, graph[i][parent[i]])
	}
}

func minKey(key []int, mstSet []bool, vertexCount int) int {
	min := int(^uint(0) >> 1)
	minIndex := -1

	for v := 0; v < vertexCount; v++ {
		if mstSet[v] == false && key[v] < min {
			min = key[v]
			minIndex = v
		}
	}

	return minIndex
}

func main() {
	graph := [][]int{
		{0, 2, 0, 6, 0},
		{2, 0, 3, 8, 5},
		{0, 3, 0, 0, 7},
		{6, 8, 0, 0, 9},
		{0, 5, 7, 9, 0},
	}

	PrimMST(graph)
}