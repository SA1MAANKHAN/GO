package main

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
}

func (b *bst) insert(val int) {
	if b.root.value == 0 {
		b.root.value = val
	} else if b.root.value <= val {
		b.root.left.value = val
	} else {
		b.root.right.value = val
	}
}

func (b *bst) display() {

}

func init() {
	tree := bst{
		root: nil,
	}
}

func main() {
	tree.insert(34)
}
