package main

import "fmt"

var m map[string]int

func main() {
	m = map[string]int{
		"Salmaan": 19,
		"Mohit":   20,
		"Anmol":   19,
		"Saurav":  15,
	}
	fmt.Println(m)
	fmt.Printf("%T", m)
}
