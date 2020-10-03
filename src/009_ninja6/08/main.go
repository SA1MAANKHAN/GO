package main

import "fmt"

func main() {
	x := foo()
	n := x()
	fmt.Println(n)
}
func foo() func() string  {
	fmt.Println("this function was assigned to a variable")
	return func() string {
		return ("hello")
		// d
	}
}
