package main

import "fmt"

func main() {
	defer bar(4)
	foo(5)
}
func foo(x int) {
	fmt.Println("you called foo , here is your number", x)
}
func bar(x int) {
	fmt.Println("you called bar , here is your number", x)
}
