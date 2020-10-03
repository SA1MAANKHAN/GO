package main

import "fmt"

var i int
var a = 8
var b = 12

func main() {
	for i = 1; i <= 16; i++ {
		if i < a && i < b {
			fmt.Printf("%v is smaller than both\n", i)
		} else if i >= a && i <= b {
			fmt.Printf("%v is in middle\n", i)
		} else {
			fmt.Printf("%v is bigger than both\n", i)
		}
	}
}
