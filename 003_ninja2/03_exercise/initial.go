package main

import "fmt"

func main() {

	const hello = "hello , untyped"
	const typedhello string = "hello , typed"

	fmt.Print("\n", hello)
	fmt.Printf("%T\n", hello)
	fmt.Print("\n", typedhello)
	fmt.Printf("%T\n", typedhello)
}
