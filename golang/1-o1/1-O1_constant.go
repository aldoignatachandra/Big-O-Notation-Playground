/*
   O1_constant.go
   Big-O: O(1) – constant time
   We do the same amount of work regardless of N.
*/

package o1

import (
	"fmt"
	"time"
)

func GetMiddleElement(arr []int) int {
	// Single arithmetic operation + array access ⇒ constant
	return arr[len(arr)/2]
}

// Demo runs the O(1) demonstration
func Demo() {
	N := 10000000
	big := make([]int, N)
	for i := 0; i < N; i++ {
		big[i] = i
	}

	t0 := time.Now()
	mid := GetMiddleElement(big)
	t1 := time.Now()

	fmt.Printf("middle element: %d\n", mid)
	fmt.Printf("O(1) demo took %.3f ms\n", float64(t1.Sub(t0).Microseconds())/1000)
}
