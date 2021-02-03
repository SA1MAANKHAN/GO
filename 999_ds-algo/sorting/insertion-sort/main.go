package main

import "fmt"

// not working yet
func main() {
	array := [7]int{6, 2, 9, 4, 7, 1}
	fmt.Println("orignal array")
	fmt.Println(array)
	fmt.Println("-----------------")
	i := 0
	for i < 7 {
		for j := 0; j < i+1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
			fmt.Println(array)
		}
		i++
	}
}
