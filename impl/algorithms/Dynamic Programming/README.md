# Dynamic Programming Examples in Go

This repository contains three well-commented programs in Go that demonstrate the concept of dynamic programming.

## What is Dynamic Programming?

Dynamic programming is a problem-solving technique used to solve complex problems by breaking them down into overlapping subproblems. It involves solving each subproblem only once and storing the solution for future use, which can greatly improve the efficiency of computations.

Dynamic programming is especially useful for optimization problems and combinatorial problems, where the solution can be expressed as a combination of solutions to smaller subproblems.

## Programs

1. **Fibonacci Sequence using Dynamic Programming:** This program efficiently computes the Fibonacci sequence using a memoization table to store previously computed values. [Source code](fibonacci.go)

2. **Longest Common Subsequence (LCS) using Dynamic Programming:** This program finds the length of the longest common subsequence between two strings using dynamic programming and a memoization table. [Source code](lcs.go)

3. **0/1 Knapsack Problem using Dynamic Programming:** This program solves the 0/1 Knapsack problem, maximizing the value of items that can be included in a knapsack with limited capacity, using dynamic programming and a memoization table. [Source code](knapsack.go)

Feel free to explore the source code of each program for detailed comments and explanations.

## Getting Started

To run these programs, make sure you have Go installed on your machine. Clone this repository and navigate to the program's directory you want to execute. Then, run the following command:

```shell
go run <program_name>.go
```

Replace <program_name> with the name of the program