// its not complete
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}
func receive(c <-chan int, q <-chan int) {
	select {
	case <-c:
		for v := range c {
			fmt.Println(v)
		}
	case <-q:
		return
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	q <- 1
	return c
}
