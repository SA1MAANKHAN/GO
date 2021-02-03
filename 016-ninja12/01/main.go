package main

import "fmt"

//Age in dog years
var Age int

func main() {
	fmt.Println("Enter the age of your dog")
	fmt.Scan(&Age)
	dog.dogAge(Age)
}

//
// func dogAge(n int) {
// 	fmt.Printf("if your dog is -- %v -- years old , it is equal to -- %v -- human years", n, n*7)
// }
