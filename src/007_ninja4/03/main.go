package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Print("\n", x[:5])
	fmt.Print("\n", x[5:])
	fmt.Print("\n", x[2:7])
	fmt.Print("\n", x[1:6])

	fmt.Printf("\n type  of x is %T\t", x)
}
