package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("goroutines:\t ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("goroutines:\t ", runtime.NumGoroutine())
	fmt.Println("count:\t\t ", counter)
	fmt.Println("Cpus:\t\t ", runtime.NumCPU())
}
