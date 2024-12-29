package main

import "fmt"

func fibonacci(n int) int {
    // Create a memoization table to store previously computed values
    memo := make([]int, n+1)
    
    // Base cases
    memo[0] = 0
    memo[1] = 1
    
    // Compute Fibonacci sequence using dynamic programming
    for i := 2; i <= n; i++ {
        memo[i] = memo[i-1] + memo[i-2]
    }
    
    return memo[n]
}

func main() {
    n := 10
    fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacci(n))
}