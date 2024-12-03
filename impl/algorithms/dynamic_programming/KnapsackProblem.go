package main

import "fmt"

type Item struct {
    value, weight int
}

func knapsack(items []Item, capacity int) int {
    n := len(items)
    
    // Create a memoization table
    memo := make([][]int, n+1)
    for i := range memo {
        memo[i] = make([]int, capacity+1)
    }
    
    // Fill the memoization table using dynamic programming
    for i := 1; i <= n; i++ {
        for w := 1; w <= capacity; w++ {
            if items[i-1].weight <= w {
                memo[i][w] = max(items[i-1].value+memo[i-1][w-items[i-1].weight], memo[i-1][w])
            } else {
                memo[i][w] = memo[i-1][w]
            }
        }
    }
    
    return memo[n][capacity]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    items := []Item{
        {value: 60, weight: 10},
        {value: 100, weight: 20},
        {value: 120, weight: 30},
    }
    capacity := 50
    fmt.Printf("Maximum value: %d\n", knapsack(items, capacity))
}