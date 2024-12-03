package main

import "fmt"

// linearSearch performs a linear search on the given slice of integers.
// It returns the index of the first occurrence of the target value, or -1
// if the target value is not found in the slice.
func linearSearch(arr []int, target int) int {
    for i, val := range arr {
        if val == target {
            return i
        }
    }
    return -1
}

func main() {
    arr := []int{5, 10, 3, 8, 6, 12, 1}
    target := 8

    // Perform a linear search on the array for the target value
    idx := linearSearch(arr, target)

    // Print the result
    if idx == -1 {
        fmt.Printf("%d was not found in the array.\n", target)
    } else {
        fmt.Printf("%d was found at index %d.\n", target, idx)
    }
}
