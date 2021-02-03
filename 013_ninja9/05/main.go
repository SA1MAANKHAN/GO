package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("initial goroutines: ", runtime.NumGoroutine())
	var counter int64
	const gos = 50
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < gos; i++ {
		go func() {
			atomic.AddInt64(&counter, 2)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("mid goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("end goroutines: ", runtime.NumGoroutine())
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("counter: ", counter)

}
