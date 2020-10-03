package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 54, 23, 624, 4, 21}
	for i := range x {
		fmt.Print("\n", x[i])
	}
	fmt.Printf("\n type  of x is %T\t", x)
}
