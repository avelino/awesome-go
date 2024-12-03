package main

import "fmt"

// The bubbleSort function takes an array of integers and sorts it in ascending order using the bubble sort algorithm
func bubbleSort(arr []int) {
    n := len(arr)

    // Traverse through all array elements
    for i := 0; i < n-1; i++ {

        // Last i elements are already in place
        for j := 0; j < n-i-1; j++ {

            // Swap the elements if they are in wrong order
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    // Example usage
    arr := []int{64, 34, 25, 12, 22, 11, 90}

    fmt.Println("Original array:", arr)

    bubbleSort(arr)

    fmt.Println("Sorted array:", arr)
}
