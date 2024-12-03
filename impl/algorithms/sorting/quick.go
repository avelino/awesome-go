package main

import "fmt"

// Function to partition the array
func partition(arr []int, low, high int) int {
    // Choose the rightmost element as pivot
    pivot := arr[high]
    // Index of smaller element and indicates the right position
    // of pivot found so far
    i := low - 1

    for j := low; j <= high-1; j++ {
        // If current element is smaller than or equal to pivot
        if arr[j] <= pivot {
            // Increment index of smaller element
            i++
            // Swap arr[i] and arr[j]
            arr[i], arr[j] = arr[j], arr[i]
        }
    }

    // Swap arr[i+1] and arr[high] (or pivot)
    arr[i+1], arr[high] = arr[high], arr[i+1]
    // Return partition index + 1
    return i + 1
}

// Function implementing QuickSort
func quickSort(arr []int, low, high int) {
    if low < high {
        // Find pivot index such that, all elements left to pivot are smaller and
        // all elements right to pivot are greater
        pi := partition(arr, low, high)
        // Recursively apply the same on left and right partitions
        quickSort(arr, low, pi-1)
        quickSort(arr, pi+1, high)
    }
}

func main() {
    // Input array
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    // Get length of the array
    n := len(arr)

    // Call quickSort function
    quickSort(arr, 0, n-1)

    // Print the sorted array
    fmt.Println("Sorted array is:", arr)
}
