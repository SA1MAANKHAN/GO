package main

import "fmt"

//this is  struct Student
type Student struct {
	name   string
	rollno int
	year   string
	group  string
}

func main() {
	s1 := Student{
		name:   "Neha Sharma",
		rollno: 11801015,
		year:   "first",
		group:  "C2",
	}
	s2 := Student{
		name:   "Raman Tyagi",
		rollno: 11801087,
		year:   "third",
		group:  "C5",
	}
	s3 := Student{
		name:   "Meyank Gupta",
		rollno: 11801014,
		year:   "first",
		group:  "C2",
	}
	Students := []Student{s1, s2, s3}
	for _, v := range Students {
		fmt.Println(v)
	}
}
