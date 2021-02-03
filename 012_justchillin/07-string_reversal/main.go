package main

import "fmt"

var inputString string
var outputString []string

func main() {
	fmt.Println("Enter the string to be reversed")
	fmt.Scanln(&inputString)
	fmt.Println("You entered: ", inputString)
	s := Reverse(inputString)
	fmt.Println(s)
}

//Reverse returns string backwards
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
