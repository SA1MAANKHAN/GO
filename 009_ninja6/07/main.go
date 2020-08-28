package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("this function was assigned to a variable")
	}
	x()
}
