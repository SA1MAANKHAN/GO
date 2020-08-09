package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(" the orignal slice \n", x)

	a := append(x, 52)
	fmt.Println(" the appended slice \n", a)

	b := append(a, 53, 54, 55)
	fmt.Println(" the twice appended slice \n", b)

	y := []int{56, 57, 58, 59, 60}

	c := append(b, y...)
	fmt.Println(" slice to slice  \n", c)
}
