package main

import "fmt"

//selection sort

func main() {
	array := [10]int{3, 5, 12, 6, 2, 67, 24, 7, 224, 4}
	fmt.Printf(" %v \n", array)

	for j := 0; j < 9; j++ {
		for i := j + 1; i <= 9; i++ {
			if array[j] > array[i] {
				// swaping
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
		fmt.Println(array)
	}

}
