package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		y := x - 64
		fmt.Print(y)

		for i := 0; i < 3; i++ {
			fmt.Printf("\n%#U", x)
		}
		fmt.Print("\n")
	}
}
