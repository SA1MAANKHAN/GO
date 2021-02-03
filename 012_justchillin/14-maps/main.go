package main

import "fmt"

func main() {
	dict := map[string]int{
		"salmaan": 20,
		"anmol":   19,
		"mohit":   20,
	}
	fmt.Println(dict)
	fmt.Printf("%T\n", dict)
	for i, v := range dict {
		fmt.Println(i, v)
	}
}
