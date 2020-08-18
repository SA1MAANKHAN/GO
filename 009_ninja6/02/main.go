package main

import "fmt"

func main() {

	s := []int{2, 4, 3, 4, 6, 4, 2, 3}
	foo(s...)
	sum(s...)
	bar(s)

}

func foo(x ...int) {
	fmt.Println("", x)
}
func sum(x ...int) {
	add := 0
	for i := 0; i < len(x); i++ {

		add = add + x[i]

	}
	fmt.Println("the sum is ", add)
}
func bar(x []int) {
	add := 0
	for _, v := range x {
		add = add + v
	}
	fmt.Println("the sum of slice is ", add)
}
