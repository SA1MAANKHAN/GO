package main

import (

	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  Dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(Dog.YearsTwo(20))
}
