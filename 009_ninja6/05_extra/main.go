package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}
func (s square) area() float64 {
	return s.side * s.side
}

func main() {
	c := circle{
		radius: 3,
	}
	s := square{
		side: 3,
	}
	fmt.Println("area of the circle is ", c.area())
	fmt.Println("area of the square is ", s.area())
}
