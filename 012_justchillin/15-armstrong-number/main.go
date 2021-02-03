package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string

// var array []int

func main() {
	fmt.Println("-----------------------------------")
	fmt.Println("WELCOME TO ARMSTRONG NUMBER CHECKER")
	fmt.Println("-----------------------------------------")
	fmt.Println("PLEASE ENTER THE NUMBER YOU WANT TO CHECK")
	fmt.Println("-----------------------------------------")
	fmt.Scanln(&input)
	newArr := strings.Split(input, "") // split on each char
	a := make([]int, 10)
	for i, v := range newArr {
		val, _ := strconv.Atoi(v)
		a[i] = val
	}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum = a[i]*a[i]*a[i] + sum
	}
	num, _ := strconv.Atoi(input)
	if num == sum {
		fmt.Println("-----------------------------------------")
		fmt.Printf("Yes! , %v is an armstrong number\n", num)
	} else {
		fmt.Println("-----------------------------------------")
		fmt.Printf("Sorry , %v is not an armstrong number\n", num)
	}
	fmt.Println("Exiting...")
}
