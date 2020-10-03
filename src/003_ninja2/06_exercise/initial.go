package main

import "fmt"

const (
	x = 2017 + iota
	y = 2017 + iota
	z = 2017 + iota
	u = 2017 + iota
)

func main() {

	fmt.Print("\n", x)
	fmt.Print("\n", y)
	fmt.Print("\n", z)
	fmt.Print("\n", u)

}
