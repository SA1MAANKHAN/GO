package main

import "fmt"

var x int
var y string
var z bool
var a string

func main() {
	x = 42
	y = "JB"
	z = true
	a = fmt.Sprintf("%v%v%v", x, y, z)
	fmt.Println(a)
	fmt.Printf("%T", a)
}
