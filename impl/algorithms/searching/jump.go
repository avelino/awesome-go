package main

import (
    "fmt"
    "math"
)

// Jump Search function
func jumpSearch(arr []int, x int) int {
    n := len(arr)
    // Finding block size to be jumped
    step := int(math.Sqrt(float64(n)))
    // Finding the block where element is present (if it is present)
    prev := 0
    for arr[int(math.Min(float64(step), float64(n))-1)] < x {
        prev = step
        step += int(math.Sqrt(float64(n)))
        if prev >= n {
            return -1
        }
    }
    // Doing a linear search for x in block beginning with prev
    for arr[prev] < x {
        prev++
        // If we reached next block or end of array, element is not present
        if prev == int(math.Min(float64(step), float64(n))) {
            return -1
        }
    }
    // If element is found
    if arr[prev] == x {
        return prev
    }
    // Element not found
    return -1
}

func main() {
    arr := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
    x := 55
    index := jumpSearch(arr, x)
    if index == -1 {
        fmt.Printf("Element %d is not present in the array\n", x)
    } else {
        fmt.Printf("Element %d is present at index %d\n", x, index)
    }
}
