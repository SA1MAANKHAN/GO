package main

import "fmt"

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

// Insert inserts new node at the end of  from linked list
func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// InsertAt inserts new node at given position
func (l *LinkedList) InsertAt(pos int, value int) {
	// create a new node
	newNode := Node{}
	newNode.value = value
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

// Print displays all the nodes from linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Print("--> Node: ", ptr.value)
		ptr = ptr.next
	}
}

// Search returns node position with given value from linked list
func (l *LinkedList) Search(val int) int {
	flag := 0
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			fmt.Println("found")
			flag = 1
			return i
		}
		ptr = ptr.next
	}
	if flag == 0 {
		fmt.Println("not found")
	}
	return -1
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func main() {
	var option int
	var data1 int
	var List LinkedList
	var k1 int

	for ok := true; ok; ok = option != 0 {

		fmt.Println("\n\nSelect the operation you want Perform")
		fmt.Println("Press 0 to quit")
		fmt.Println("Press 1 for appendnode")
		fmt.Println("Press 2 for Search")
		fmt.Println("Press 3 for insertnodeafter")
		fmt.Println("Press 4 for printlist")

		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("Exiting.....")
			break
		case 1:
			fmt.Println("Append Node Operation \n data of the Node to be Appended")
			fmt.Scanln(&data1)
			List.Insert(data1)
		case 2:
			fmt.Println("Search Node Operation \nEnter  value of the Node to be Searched")
			fmt.Scanln(&data1)
			List.Search(data1)

		case 3:
			fmt.Println("Insert Node After Operation \nEnter key of existing Node after which you want to Insert this New node: ")
			fmt.Println("Enter key & data of the New Node first: ")
			fmt.Scanln(&k1)
			fmt.Scanln(&data1)
			List.InsertAt(k1, data1)

		case 4:
			List.Print()

		default:
			fmt.Println("Enter a valid option")
		}
	}
}
