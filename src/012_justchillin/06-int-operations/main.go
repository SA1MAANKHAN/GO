package main

import "fmt"

var a int
var b int
var c float64

func main() {
	a = 9
	b = 2
	d := a / b
	c = float64(a / b)
	fmt.Println(c, d)
}
