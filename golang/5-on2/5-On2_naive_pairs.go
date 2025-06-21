/*
   On2_naive_pairs.go
   Big-O: O(n²) – quadratic
   Compare every pair of elements (e.g., brute-force duplicate search).
*/

package onlogn

import (
	"fmt"
	"time"
)

func hasDuplicate(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func Demo() {
	N := 30000 // 30k² ≈ 900M comparisons ⇒ slow
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i
	}
	nums = append(nums, N-1) // guarantee duplicate at the tail (worst-case)

	t0 := time.Now()
	dup := hasDuplicate(nums)
	t1 := time.Now()

	fmt.Printf("duplicate found: %t\n", dup)
	fmt.Printf("O(n²) demo with N=%d took %.3f ms\n", N, float64(t1.Sub(t0).Microseconds())/1000)
}
