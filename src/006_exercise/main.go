package main

import "fmt"

func main() {
	// slice
	x := []int{2, 5, 3, 6, 7}

	for i, v := range x {
		fmt.Printf("%v  %v\n", i, v)
	}
}
