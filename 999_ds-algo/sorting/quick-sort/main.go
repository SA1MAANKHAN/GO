package main

import "fmt"

// yet to complete

func main() {
	array := []int{9, 8, 5, 10, 7, 6, 3, 1, 2, 4, 11}
	for i, v := range array {
		fmt.Printf("%v [%v]\n", i, v)
	}

	pivot := array[5]
	backcounter := 10
	i := 0
	for i < 5 {
		if (array[i]) > pivot && array[backcounter] < pivot {
			swap(&array[i], &array[backcounter])
			backcounter--
		} else if (array[i] > pivot) && array[backcounter] > pivot {
			backcounter--
		} else {
			i++
		}
	}
	fmt.Println("-------------------")
	for i, v := range array {
		fmt.Printf("%v [%v]\n", i, v)
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
