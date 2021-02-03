package main

import "fmt"

var top int
var arr []int
var option int
var pos int
var value int

// functions
func isEmpty() bool {
	if top == -1 {
		return true
	} else {
		return false
	}
}
func isFull() bool {
	if top == 4 {
		return true
	} else {
		return false
	}
}
func push(value int) {
	if isFull() {
		fmt.Println("The stack is full , OVERFLOW!")
	} else {
		arr[top] = value
	}
}
func pop() int {
	var value int
	if isEmpty() {
		fmt.Println("The stack is already empty , UNDERFLOW!")
		return 0
	} else {
		arr[top] = value
		top++
		return value
	}
}

// func count() int {
// 	return top + 1
// }
func change(pos int, value int) {
	arr[pos] = value
}
func peek(pos int) int {
	if isEmpty() {
		fmt.Println("The stack is already empty , UNDERFLOW!")
		return 0
	} else {
		return arr[len(arr)-1]
	}
}
func display() {
	if isEmpty() {
		fmt.Println("The stack is already empty , UNDERFLOW!")
	} else {
		for i, v := range arr {
			fmt.Printf("postion: %v ---- value: %v\n", i, v)
		}
	}
}

func main() {

	top = -1
	arr = []int{0, 0, 0, 0, 0}
	for ok := true; ok; ok = option != 0 {

		fmt.Println("Select the operation you want Perform . Enter 0 for exit.")
		fmt.Println("Press 1 for Push")
		fmt.Println("Press 2 for Pop")
		fmt.Println("Press 3 for isEmpty")
		fmt.Println("Press 4 for isFull")
		fmt.Println("Press 5 for peek")
		fmt.Println("Press 6 for count")
		fmt.Println("Press 7 for change or update")
		fmt.Println("Press 8 for display")

		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Enter the value you want to push")
			fmt.Scanln(&value)
			push(value)

		case 2:
			fmt.Printf("Pop function called - poped value: %v", pop())

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
			fmt.Println("Enter the postion of element you want to pick")
			fmt.Scanln(&pos)
			fmt.Printf("Peek funtion called - value at postion %v : %v ", pos, peek(pos))
			//
		// case 6:
		// 	fmt.Printf("Number of elemets :%v ", count())

		case 7:
			fmt.Println("change funtion called ")
			fmt.Println("Enter the postion of item you want to change")
			fmt.Scanln(&pos)
			fmt.Println("Enter the value of the item you want to change")
			fmt.Scanln(&value)
			change(pos, value)
			fmt.Printf("Peek funtion called - value at postion %v : %v ", pos, peek(pos))

		case 8:
			fmt.Println("display funtion called")
			display()

		}
	}
}
