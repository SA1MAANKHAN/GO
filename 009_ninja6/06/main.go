package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is an anonymus function")
	}()
	fmt.Println("This is a simpe statement")
}
