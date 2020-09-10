package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 3
	fmt.Println(<-c)
}

// this does not work
