package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p person) speak() {
	fmt.Println("i am", p.fname, p.lname, "and i am ", p.age)
}

func main() {
	p1 := person{
		fname: "Salmaan",
		lname: "Khan",
		age:   19}
	p1.speak()
}
