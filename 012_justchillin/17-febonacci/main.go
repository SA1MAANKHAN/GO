package main

import "fmt"

var input int

func main() {
	fmt.Println("Enter the number you want to find febonacci upto : ")
	fmt.Scan(&input)
	febonacci(input)
	fmt.Println(febonacciRecursive(input))

}
func febonacci(x int) {
	array := make([]int, x+1)
	array[0] = 0
	array[1] = 1
	for i := 2; i < x+1; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	fmt.Println(array[x])
}

func febonacciRecursive(x int) int {
	if x < 2 {
		return x
	}
	return febonacciRecursive(x-1) + febonacciRecursive(x-2)
}
