package main

import (
    "fmt"
    "math"
)

func lcsLength(s1, s2 string) int {
    m := len(s1)
    n := len(s2)
    
    // Create a memoization table
    memo := make([][]int, m+1)
    for i := range memo {
        memo[i] = make([]int, n+1)
    }
    
    // Fill the memoization table using dynamic programming
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s1[i-1] == s2[j-1] {
                memo[i][j] = memo[i-1][j-1] + 1
            } else {
                memo[i][j] = int(math.Max(float64(memo[i-1][j]), float64(memo[i][j-1])))
            }
        }
    }
    
    return memo[m][n]
}

func main() {
    s1 := "AGGTAB"
    s2 := "GXTXAYB"
    fmt.Printf("Longest Common Subsequence: %d\n", lcsLength(s1, s2))
}