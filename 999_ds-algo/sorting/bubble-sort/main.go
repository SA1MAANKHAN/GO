package main

import "fmt"

//bubble sort

func main() {
	array := [10]int{3, 5, 12, 6, 2, 67, 24, 7, 224, 4}
	fmt.Printf("%T %v \n", array, array)

	// sorted := make([]int, 10)
	for j := 0; j < 9; j++ {
		for i := 0; i <= 8; i++ {
			if array[i] > array[i+1] {
				// swaping
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
			}
		}
		fmt.Println(array)
	}

}
