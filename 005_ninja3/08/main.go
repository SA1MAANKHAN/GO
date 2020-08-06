package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Print("its not printed")
	case true:
		fmt.Print("its printed")
	}
}
