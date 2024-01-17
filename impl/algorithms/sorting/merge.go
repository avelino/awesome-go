package main

import "fmt"

// Merge two subarrays of 'arr[]'
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]
func merge(arr []int, l, m, r int) {
	// Create temporary arrays to store the two subarrays
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// Merge the two temporary arrays back into 'arr[]'

	// Initial indexes of the first and second subarrays
	i := 0 // index of first subarray
	j := 0 // index of second subarray

	// Initial index of merged subarray
	k := l

	for i < n1 && j < n2 {
		// Compare the elements from the two subarrays
		if L[i] <= R[j] {
			// The element from the left subarray is smaller or equal
			// Copy it to arr[k] and increment i
			arr[k] = L[i]
			i++
		} else {
			// The element from the right subarray is smaller
			// Copy it to arr[k] and increment j
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[], if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[], if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

// Main function that sorts arr[l..r] using merge()
func mergeSort(arr []int, l, r int) {
	if l < r {
		// Same as (l+r)/2, but avoids overflow for
		// large l and h
		m := l + (r-l)/2

		// Sort first and second halves
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		// Merge the sorted halves
		merge(arr, l, m, r)
	}
}

func main() {
	// Example usage of Merge Sort

	arr := []int{12, 11, 13, 5, 6, 7}
	n := len(arr)

	fmt.Println("Original array:", arr)

	mergeSort(arr, 0, n-1)

	fmt.Println("Sorted array:", arr)
}
