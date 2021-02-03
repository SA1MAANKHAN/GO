package main

import "fmt"

var front int
var rear int
var arr []int
var val int
var option int

func isEmpty() bool {
	if front == -1 && rear == -1 {
		return true
	} else {
		return false
	}
}
func isFull() bool {
	if rear == 4 {
		return true
	} else {
		return false
	}
}
func enqueue(value int) {
	if isFull() {
		fmt.Println("The queue is Full")
	} else if isEmpty() {
		front = 0
		rear = 0
		arr[rear] = value
	} else {
		rear++
		arr[rear] = value
	}
}
func dequeue() int {
	if isEmpty() {

		fmt.Println("Queue is already Empty")
		return 0
	} else if front == rear {
		x := arr[front]
		arr[front] = 0
		front = -1
		rear = -1
		return x
	} else {
		x := arr[front]
		arr[front] = 0
		front++
		return x
	}

}
func count() int {
	return (rear - front) + 1
}
func display() {
	for _, value := range arr {
		fmt.Println(value)
	}
}

func main() {
	front = -1
	rear = -1
	arr = []int{0, 0, 0, 0, 0}

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
			fmt.Scanln(&val)
			enqueue(val)

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
