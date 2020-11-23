package main

import "fmt"

func main() {
	fmt.Println(add(4, 4, 3, 2))
}

func add(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum

}
