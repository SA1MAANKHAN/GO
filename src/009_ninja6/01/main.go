package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(bar())
	fmt.Println("this is foo() ", s)
}

func foo() int {
	x := 4
	return x + 1
}
func bar() (int, string) {
	y := 5
	return y + 1, "this is bar()"
}
