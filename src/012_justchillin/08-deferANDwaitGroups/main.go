package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// printa()
	// defer printa()
	// defer printb()
	// defer printc()
	go printa()
	wg.Add(1)
	go printb()
	wg.Add(1)
	go printc()
	wg.Add(1)
	wg.Wait()

}

func printa() {
	fmt.Println("a")
	fmt.Println("a")
	fmt.Println("a")
	fmt.Println("a")
	fmt.Println("a")
	fmt.Println("a")
	fmt.Println("a")
	fmt.Println("a")

}

func printb() {
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
	fmt.Println("b")
}

func printc() {
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
	fmt.Println("c")
}
