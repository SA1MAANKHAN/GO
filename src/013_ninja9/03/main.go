package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("initial goroutines: ", runtime.NumGoroutine())
	var counter int
	const gos = 50
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < gos; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			fmt.Println("counter ", counter)
		}()
		go func() {
			u := counter
			u++
			counter = u
			fmt.Println("counter ", counter)
			wg.Done()
		}()
		fmt.Println("mid goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("end goroutines: ", runtime.NumGoroutine())
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("counter: ", counter)

}
