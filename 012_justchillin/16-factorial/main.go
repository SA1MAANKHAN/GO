package main

import "fmt"

var input int

func main() {
	fmt.Println("Enter the number you want to find factorial of : ")
	fmt.Scan(&input)
	factorialIterative(input)
	fmt.Println("Recursive factorial is -- > ", factorialRecursive(input))
}

func factorialIterative(x int) {

	factorial := 1
	for i := 1; i <= x; i++ {
		factorial = factorial * i
	}
	fmt.Println("Iterative factorial is -->  ", factorial)

}

func factorialRecursive(x int) int {
	if x == 2 {
		return 2
	}
	return x * factorialRecursive(x-1)

}
