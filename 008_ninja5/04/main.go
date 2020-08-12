package main

import "fmt"

func main() {
	s := struct {
		mitar   []string
		contact map[string]int
	}{
		mitar:   []string{"lav", "ans", "anm", "jat", "abv"},
		contact: map[string]int{"cad": 7, "ans": 17, "anm": 9, "jat": 2, "abv": 2},
	}
	fmt.Println(s)
}
