package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}
	for i, v := range x {
		fmt.Print("\n", i, v)
	}
	fmt.Printf("\ntype of x is %T\t", x)
}
