/*
   OnLogN_merge_sort.go
   Big-O: O(n log n)
   Classic merge-sort: log n levels of splitting × n work each level.
*/

package on2

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) >> 1
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(a, b []float64) []float64 {
	result := make([]float64, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}

func Demo() {
	N := 200000 // use smaller N, merge-sort is heavy
	nums := make([]float64, N)
	for i := 0; i < N; i++ {
		nums[i] = rand.Float64()
	}

	t0 := time.Now()
	mergeSort(nums)
	t1 := time.Now()

	fmt.Printf("O(n log n) demo with N=%d took %.3f ms\n", N, float64(t1.Sub(t0).Microseconds())/1000)
}
