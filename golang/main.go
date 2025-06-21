package main

import (
	"fmt"
	"os"

	o1 "./1-o1"
	ologn "./2-ologn"
	on "./3-on"
	onlogn "./4-onlogn"
	on2 "./5-on2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify which demo to run: o1, ologn, on, on2, onlogn or all")
		return
	}

	demo := os.Args[1]

	switch demo {
	case "o1":
		o1.Demo()
	case "ologn":
		ologn.Demo()
	case "on":
		on.Demo()
	case "onlogn":
		onlogn.Demo()
	case "on2":
		on2.Demo()
	case "all":
		fmt.Println("\n--- O(1) Constant Time ---")
		o1.Demo()
		fmt.Println("\n--- O(log n) Logarithmic Time ---")
		ologn.Demo()
		fmt.Println("\n--- O(n) Linear Time ---")
		on.Demo()
		fmt.Println("\n--- O(n log n) Time ---")
		onlogn.Demo()
		fmt.Println("\n--- O(n²) Quadratic Time ---")
		on2.Demo()
	default:
		fmt.Println("Unknown demo type. Please use: o1, ologn, on, onlogn, on2, or all")
	}
}
