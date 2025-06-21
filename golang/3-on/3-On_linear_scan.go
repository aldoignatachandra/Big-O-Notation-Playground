/*
   On_linear_scan.go
   Big-O: O(n) – linear
   Iterates once over the entire collection (sum).
*/

package on

import (
	"fmt"
	"time"
)

func sum(arr []int) int {
	total := 0
	for _, n := range arr {
		total += n
	}
	return total
}

func Demo() {
	N := 10000000
	big := make([]int, N)
	for i := 0; i < N; i++ {
		big[i] = i
	}

	t0 := time.Now()
	total := sum(big)
	t1 := time.Now()

	fmt.Printf("sum: %d\n", total)
	fmt.Printf("O(n) demo with N=%d took %.3f ms\n", N, float64(t1.Sub(t0).Microseconds())/1000)
}
