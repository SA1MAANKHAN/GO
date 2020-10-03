package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	c := make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			mu.Lock()
			for i := 0; i < 10; i++ {
				c <- i + 1
			}
			mu.Unlock()
		}()
	}
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

}
