/*
   OlogN_binary_search.go
   Big-O: O(log n) – logarithmic
   Binary search halves the remaining search space each step.
*/

package ologn

import (
	"fmt"
	"time"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) >> 1 // use bitwise right shift to divide by 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func Demo() {
	N := 10000000
	bigSorted := make([]int, N)
	for i := 0; i < N; i++ {
		bigSorted[i] = i
	}

	t0 := time.Now()
	idx := binarySearch(bigSorted, N-1) // worst-case (near end)
	t1 := time.Now()

	fmt.Printf("found at index: %d\n", idx)
	fmt.Printf("O(log n) demo with N=%d took %.3f ms\n", N, float64(t1.Sub(t0).Microseconds())/1000)
}
