package main

import (
	"fmt"
)

type coustomErr struct {
	s string
}

func (ce coustomErr) Error() string {
	return fmt.Sprintf("this is the coustom error %v", ce.s)
}

func main() {
	c1 := coustomErr{
		s: "i am just a string ",
	}
	foo(c1)
}

func foo(err error) {
	fmt.Println("foo ran -", err)
	// if err != nil {
	// 	log.Println(err)
	// }
}
