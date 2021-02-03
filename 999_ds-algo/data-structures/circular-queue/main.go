package main

import "fmt"

var front int
var rear int
var arr []int
var itemcounter int
var option int
var value int

func isEmpty() bool {
	if rear == -1 && front == -1 {
		return true
	} else {
		return false
	}
}

func isFull() bool {
	if (rear+1)%5 == front {
		return true
	} else {
		return false
	}
}
func enqueue(value int) {
	if isFull() {
		fmt.Println("Queue is already full")
		return
	} else if isEmpty() {
		front = 0
		rear = 0
	} else {
		rear = (rear + 1) % 5
	}
	arr[rear] = value
	itemcounter++
}

func dequeue() int {
	if isEmpty() {
		fmt.Println("Queue is already empty")
		return 0
	} else if front == rear {
		x := arr[front]
		arr[front] = 0
		front = -1
		rear = -1
		itemcounter--
		return x
	} else {
		x := arr[front]
		arr[front] = 0
		front = (front + 1) % 5
		itemcounter--
		return x
	}
}
func count() int {
	return itemcounter
}
func display() {
	for i, v := range arr {
		fmt.Printf("\n%v-----%v", i, v)
	}

}
func main() {

	front = -1
	rear = -1
	arr = []int{0, 0, 0, 0, 0}
	itemcounter = 0

	for ok := true; ok; ok = option != 0 {

		fmt.Println("\n\nSelect the operation you want Perform")
		fmt.Println("Press 0 to quit")
		fmt.Println("Press 1 for enqueue")
		fmt.Println("Press 2 for dequeue")
		fmt.Println("Press 3 for isEmpty")
		fmt.Println("Press 4 for isFull")
		fmt.Println("Press 5 for count")
		fmt.Println("Press 6 for display")

		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("Bye!")
			break
		case 1:
			fmt.Println("Enter the value you want to Enqueue")
			fmt.Scanln(&value)
			enqueue(value)

		case 2:
			fmt.Println("Dequeue Selected -----Removed item is :")
			fmt.Println(dequeue())

		case 3:
			if isEmpty() {
				fmt.Println("Stack is Empty")
			} else {
				fmt.Println("Stack is not Empty")
			}

		case 4:
			if isFull() {
				fmt.Println("Stack is full")
			} else {
				fmt.Println("Stack is not full")
			}

		case 5:
			fmt.Printf("Total number of elements in queue are : %v ", count())

		case 6:
			for i, v := range arr {
				fmt.Printf("\nat %v ------%v ", i, v)
			}

		default:
			fmt.Println("Enter a valid option!")

		}
	}
}
