// C++ program for Dijkstra's single source shortest path
// algorithm. The program is for adjacency matrix
// representation of the graph
#include <iostream>
using namespace std;
#include <limits.h>
  
// Number of verticepackage main

import (
    "fmt"
    "math"
)

// Number of vertices in the graph
const V = 9

// A utility function to find the vertex with minimum
// distance value, from the set of vertices not yet included
// in shortest path tree
func minDistance(dist []int, sptSet []bool) int {
    // Initialize min value
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

// A utility function to print the constructed distance
// array
func printSolution(dist []int) {
    fmt.Println("Vertex \t Distance from Source")
    for i := 0; i < V; i++ {
        fmt.Printf("%d \t\t\t\t %d\n", i, dist[i])
    }
}

// Function that implements Dijkstra's single source
// shortest path algorithm for a graph represented using
// adjacency matrix representation
func dijkstra(graph [V][V]int, src int) {
    dist := make([]int, V)
    sptSet := make([]bool, V)

    // Initialize all distances as INFINITE and stpSet[] as false
    for i := 0; i < V; i++ {
        dist[i] = math.MaxInt32
        sptSet[i] = false
    }

    // Distance of source vertex from itself is always 0
    dist[src] = 0

    // Find shortest path for all vertices
    for count := 0; count < V-1; count++ {
        // Pick the minimum distance vertex from the set of
        // vertices not yet processed. u is always equal to
        // src in the first iteration.
        u := minDistance(dist, sptSet)

        // Mark the picked vertex as processed
        sptSet[u] = true

        // Update dist value of the adjacent vertices of the
        // picked vertex.
        for v := 0; v < V; v++ {
            // Update dist[v] only if is not in sptSet,
            // there is an edge from u to v, and total
            // weight of path from src to  v through u is
            // smaller than current value of dist[v]
            if !sptSet[v] && graph[u][v] != 0 && dist[u] != math.MaxInt32 && dist[u]+graph[u][v] < dist[v] {
                dist[v] = dist[u] + graph[u][v]
            }
        }
    }

    // print the constructed distance array
    printSolution(dist)
}


func main() {
    // Create a sample adjacency matrix graph
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

    // Set the source vertex for the algorithm
    src := 0

    // Call the Dijkstra's algorithm function with the graph and source vertex
    dijkstra(graph, src)
}s in the graph
#define V 9
  
// A utility function to find the vertex with minimum
// distance value, from the set of vertices not yet included
// in shortest path tree
int minDistance(int dist[], bool sptSet[]) {
  
    // Initialize min value
    int min = INT_MAX, min_index;
  
    for (int v = 0; v < V; v++)
        if (sptSet[v] == false && dist[v] <= min)
            min = dist[v], min_index = v;
  
    return min_index;
}
  
// A utility function to print the constructed distance
// array
void printSolution(int dist[]) {
    cout << "Vertex \t Distance from Source" << endl;
    for (int i = 0; i < V; i++)
        cout << i << " \t\t\t\t" << dist[i] << endl;
}
  
// Function that implements Dijkstra's single source
// shortest path algorithm for a graph represented using
// adjacency matrix representation
void dijkstra(int graph[V][V], int src) {
    int dist[V]; // The output array.  dist[i] will hold the
                 // shortest
    // distance from src to i
  
    bool sptSet[V]; // sptSet[i] will be true if vertex i is
                    // included in shortest
    // path tree or shortest distance from src to i is
    // finalized
  
    // Initialize all distances as INFINITE and stpSet[] as
    // false
    for (int i = 0; i < V; i++)
        dist[i] = INT_MAX, sptSet[i] = false;
  
    // Distance of source vertex from itself is always 0
    dist[src] = 0;
  
    // Find shortest path for all vertices
    for (int count = 0; count < V - 1; count++) {
        // Pick the minimum distance vertex from the set of
        // vertices not yet processed. u is always equal to
        // src in the first iteration.
        int u = minDistance(dist, sptSet);
  
        // Mark the picked vertex as processed
        sptSet[u] = true;
  
        // Update dist value of the adjacent vertices of the
        // picked vertex.
        for (int v = 0; v < V; v++)
  
            // Update dist[v] only if is not in sptSet,
            // there is an edge from u to v, and total
            // weight of path from src to  v through u is
            // smaller than current value of dist[v]
            if (!sptSet[v] && graph[u][v]
                && dist[u] != INT_MAX
                && dist[u] + graph[u][v] < dist[v])
                dist[v] = dist[u] + graph[u][v];
    }
  
    // print the constructed distance array
    printSolution(dist);
}
  
// driver's code
int main() {
  
    /* Let us create the example graph discussed above */
    int graph[V][V] = { { 0, 4, 0, 0, 0, 0, 0, 8, 0 },
                        { 4, 0, 8, 0, 0, 0, 0, 11, 0 },
                        { 0, 8, 0, 7, 0, 4, 0, 0, 2 },
                        { 0, 0, 7, 0, 9, 14, 0, 0, 0 },
                        { 0, 0, 0, 9, 0, 10, 0, 0, 0 },
                        { 0, 0, 4, 14, 10, 0, 2, 0, 0 },
                        { 0, 0, 0, 0, 0, 2, 0, 1, 6 },
                        { 8, 11, 0, 0, 0, 0, 1, 0, 7 },
                        { 0, 0, 2, 0, 0, 0, 6, 7, 0 } 
    };
  
    // Function call
    dijkstra(graph, 0);
  
    return 0;

}