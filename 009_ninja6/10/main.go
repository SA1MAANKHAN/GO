package main

import "fmt"

func main() {
	a := 0
	fmt.Println("Enter a value")
	fmt.Scanln(&a)
	loo(a)
}

func loo(n int) {

	b := n + 1
	fmt.Printf(" %v + 1 =  %v", n, b)
}
