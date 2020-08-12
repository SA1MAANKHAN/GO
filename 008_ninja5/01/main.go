package main

import "fmt"

type Person struct {
	first_name string
	last_name  string
	fav_flavs  []string
}

func main() {
	p1 := Person{
		first_name: "Salmaan",
		last_name:  "Khan",
		fav_flavs:  []string{"belgian chocochip", "straberry", "choco"},
	}
	p2 := Person{
		first_name: "Monty",
		last_name:  " ",
		fav_flavs:  []string{"chocochip", "dry fruit", "vanilla"},
	}
	// 	fmt.Printf("Name is %v %v and i like %v\n", p1.first_name, p1.last_name, p1.fav_icecream)
	// 	fmt.Printf("Name is %v %v and i like %v", p2.first_name, p2.last_name, p2.fav_icecream)
	//
	fmt.Println(p1, p2)
}
