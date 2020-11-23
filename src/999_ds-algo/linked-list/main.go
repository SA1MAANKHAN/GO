package main

import "fmt"

var option int
var head *node

// node is a custom type
type node struct {
	key  int
	data int
	next *node
}
type singlylinkedlist struct {
	head *node
}

func (n *node) Node(k int, d int) {
	n.key = k
	n.data = d
	n.next = nil
}

func (s *singlylinkedlist) nodeExists(k int) *node {
	var temp *node
	ptr := s.head
	for ptr != nil {
		if ptr.key == k {
			temp = ptr
		}
		ptr = ptr.next
	}
	return temp
}

func (s *singlylinkedlist) appendnode(n *node) {
	if s.nodeExists(n.key) != nil {
		fmt.Println("node already exists")
	} else if head == nil {
		head = n
		fmt.Println("Node appended")
	} else {
		ptr := head.next
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = n
		fmt.Println("Node appended")
	}
}

func (s *singlylinkedlist) prependnode(n *node) {
	if s.nodeExists(n.key) != nil {
		fmt.Println("node already exusts")
	} else {
		n.next = head
		head = n
	}
}
func (s *singlylinkedlist) insertnodeafter(k int, n *node) {
	ptr := s.nodeExists(k)
	if ptr == nil {
		fmt.Println("no node exists with that key value")
	} else {
		if s.nodeExists(n.key) != nil {
			fmt.Println("node already exists with that key value")
		} else {
			n.next = ptr.next
			ptr.next = nil
			fmt.Println("node inserted")
		}
	}
}

func (s *singlylinkedlist) deletenode(k int) {
	var temp *node
	if head == nil {
		fmt.Println("list already emplty")
	} else if head != nil {
		if head.key == k {
			head = head.next
			fmt.Println("node unlinked with key values")
		} else {
			prevptr := head
			currentptr := head.next
			for currentptr != nil {
				if currentptr.key == k {
					temp = currentptr
					currentptr = nil
				} else {
					prevptr = prevptr.next
					currentptr = currentptr.next
				}
			}
			if temp != nil {
				prevptr.next = temp.next
				fmt.Println("node unlinled with given key")
			} else {
				fmt.Println("node doesnt exist with the key")
			}
		}
	}
}
func (s *singlylinkedlist) updatenode(k int, data int) {
	ptr := s.nodeExists(k)
	if ptr != nil {
		ptr.data = data
		fmt.Println("node data updated")
	} else {
		fmt.Println("node doesnt exists with key")
	}
}
func (s *singlylinkedlist) printlist() {

	if head == nil {
		fmt.Println("no nodes in the list")
	} else {
		fmt.Println("singly linked list values")
		temp := head
		for temp != nil {
			fmt.Printf("( %v , %v) -->", temp.key, temp.data)
			temp = temp.next
		}
	}
}

func main() {
	var n1 node
	var s singlylinkedlist
	var key1 int
	var k1 int
	var data1 int

	for ok := true; ok; ok = option != 0 {

		fmt.Println("\n\nSelect the operation you want Perform")
		fmt.Println("Press 0 to quit")
		fmt.Println("Press 1 for appendnode")
		fmt.Println("Press 2 for prependnode")
		fmt.Println("Press 3 for insertnodeafter")
		fmt.Println("Press 4 for deletenodebykey")
		fmt.Println("Press 5 for updatenodebykey")
		fmt.Println("Press 6 for printlist")

		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("Existing.....")
			break
		case 1:
			fmt.Println("Append Node Operation \nEnter key & data of the Node to be Appended")
			fmt.Scanln(&key1)
			fmt.Scanln(&data1)
			n1.key = key1
			n1.data = data1
			s.appendnode(&n1)
		case 2:
			fmt.Println("Prepend Node Operation \nEnter key & data of the Node to be Prepended")
			fmt.Scanln(&key1)
			fmt.Scanln(&data1)
			n1.key = key1
			n1.data = data1
			s.prependnode(&n1)

		case 3:
			fmt.Println("Insert Node After Operation \nEnter key of existing Node after which you want to Insert this New node: ")
			fmt.Scanln(&k1)
			fmt.Println("Enter key & data of the New Node first: ")
			fmt.Scanln(&key1)
			fmt.Scanln(&data1)
			n1.key = key1
			n1.data = data1

		case 4:
			fmt.Println("Delete Node By Key Operation - \nEnter key of the Node to be deleted: ")
			fmt.Scanln(&k1)
			s.deletenode(k1)

		case 5:
			fmt.Println("Update Node By Key Operation - \nEnter key & NEW data to be updated")
			fmt.Scanln(key1)
			fmt.Scanln(data1)
			s.updatenode(key1, data1)

		case 6:
			s.printlist()

		default:
			fmt.Println("Enter a valid option")
		}
	}
}
