package main

import "fmt"

var s1, s2, s3 int

func main() {
	fmt.Println("Enter 3 sides of the triangle: ")
	fmt.Scanln(&s1, &s2, &s3)
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
	if s1 == s2 && s2 == s3 {
		fmt.Println("Its an equilateral triangle")
	} else if s1 == s2 || s2 == s3 || s1 == s3 {
		fmt.Println("Its an isoceles triangle")
	} else {
		fmt.Println("Its an scalen triangle")
	}

}
