package main

import "fmt"

func main() {
	x := 99

	a := 42 == 42
	b := 42 <= 42
	c := 42 >= 42
	d := 42 != 42
	e := 42 < 42
	f := 42 > 42
	fmt.Print("\n", x)
	fmt.Print("\n", a)
	fmt.Print("\n", b)
	fmt.Print("\n", c)
	fmt.Print("\n", d)
	fmt.Print("\n", e)
	fmt.Print("\n", f)
}
