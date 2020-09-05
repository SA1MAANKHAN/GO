package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Arch :\t", runtime.GOARCH)
	fmt.Println("OS : \t", runtime.GOOS)
}
