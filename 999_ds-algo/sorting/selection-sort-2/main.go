package main

import "fmt"

// not working yet
func main() {
	array := [10]int{6, 3, 7, 8, 2, 5, 10, 4, 1, 9}
	fmt.Println(array)

	for i := 0; i <= 8; i++ {
		min := i
		for j := i; j <= 8; j++ {
			if array[j] > array[j+1] {
				min = j + 1
			}
		}
		temp := array[i]
		array[i] = array[min]
		array[min] = temp
		fmt.Println(array)
	}
}
