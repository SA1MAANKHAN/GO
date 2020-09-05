package main

import "fmt"

var arr [4][4]int

func main() {
	arr[2][1] = 42
	fmt.Println(arr)
	fmt.Printf("%T", arr)
}
