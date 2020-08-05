package main

import "fmt"

func main() {
	x := 99

	fmt.Print("\n", x)
	fmt.Printf("%d\n", x)
	fmt.Printf("%#b\n", x)
	fmt.Printf("%#x\n", x)

y := x << 1
fmt.Print("\n", y)
	fmt.Printf("%d\n", y)
	fmt.Printf("%#b\n",y )
	fmt.Printf("%#x\n", y)
}
