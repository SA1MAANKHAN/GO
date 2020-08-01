package main

import "fmt"

type mytype int

var y int
var x mytype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 99
	fmt.Println("value of x is", x)
	y = int(x)
	fmt.Println("value of y is", y)
	fmt.Printf("%T\n", y)
}
