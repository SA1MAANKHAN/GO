package main

import "fmt"

func main() {
	x := 2000
	fmt.Print("I am alive....")
	for {

		fmt.Print("\n", x)
		x++
		if x > 2020 {
			break
		}
	}
	// fmt.Print("\nand countung.....")
}
