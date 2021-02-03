package main

// yet to complete
func main() {

}

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	length := float64(len(array))
	middle := (int)(length / 2)
	left := array[1:middle]
	right := array[middle+1:]

	return merge(left , right)
}

func merge(left []int, right []int) []int{
	result := [10]
}
