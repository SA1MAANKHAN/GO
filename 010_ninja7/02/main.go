package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	x := person{
		firstName: "Mera Naam",
		lastName:  "chin chin chin",
		age:       10,
	}
	fmt.Println(x.firstName, x.lastName, x.age)
	changeme(&x)
	fmt.Println(x.firstName, x.lastName, x.age)
	// fmt.Println((*x).firstName, (*x).lastName, (*x).age)

}

func changeme(x *person) {

	x.firstName = "Salmaan"
	x.lastName = "Khan"
	x.age = 19

}
