package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const gos = 50
	counter := 0
	counter2 := 2
	var mu sync.Mutex

	wg.Add(gos)
	for i := 0; i < gos; i++ {
		go func() {
			mu.Lock()
			v := counter
			v += 5
			counter = v
			mu.Unlock()
		}()
		go func() {
			mu.Lock()
			fmt.Println("this is the counter", counter)
			mu.Unlock()

		}()
		wg.Done()

	}
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	fmt.Println("this is the second counter ", counter2)
	wg.Wait()

}
