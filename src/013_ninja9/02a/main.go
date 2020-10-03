// still didnt got that completely but thats ok!
package main

import "fmt"

type person struct {
	fname string
	// lived int
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("oktan")
}
func saysomething(h human) {
	h.speak()
}

func main() {
	p := person{"Oktan"}
	saysomething(&p)
}
