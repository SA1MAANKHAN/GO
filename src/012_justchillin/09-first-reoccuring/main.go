package main

import "fmt"

var arr []int
var check []int
var size int
var input int

func main() {
	fmt.Println("Enter the size of the array")
	arr = []int{2, 5, 25, 6, 25, 55, 67, 3}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				fmt.Println(arr[i], " has repeated")
			}
		}
	}

}
