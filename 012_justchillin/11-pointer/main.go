package main

import "fmt"

//Node is custom type which has key , data and pointer to the next node
type Node struct {
	key  int
	data int
	next *Node
}

// Head points to the first node
type Head struct {
	next *Node
}

func main() {
	n1 := Node{
		5,
		2,
		nil,
	}
	head := Head{&n1}
	fmt.Println(head)
	fmt.Println(n1)
}
