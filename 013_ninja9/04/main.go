package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("initial goroutines: ", runtime.NumGoroutine())
	var counter int
	var mu sync.Mutex
	const gos = 50
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < gos; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
			fmt.Println("counter ", counter)
		}()
		// go func() {
		// 	mu.Lock()
		// 	u := counter
		// 	runtime.Gosched()
		// 	u++
		// 	counter = u
		// 	mu.Unlock()
		// 	wg.Done()
		// 	fmt.Println("counter ", counter)
		// }()
		fmt.Println("mid goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("end goroutines: ", runtime.NumGoroutine())
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("counter: ", counter)

}
