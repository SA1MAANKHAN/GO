package main

import "fmt"

func main() {
	x := 2
	fmt.Println("the address of x is ", &x)
	fmt.Printf("the value stored at %v is %v", &x, *&x)
}
