package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v divided by 4 gives remainder  %v\n", i, i%4)
	}
}
