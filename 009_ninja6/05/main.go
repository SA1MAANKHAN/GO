package main

import (
	"fmt"
	"math"
)

func main() {
	areaCircle(5.42)
	areaSquare(4.32)

}
func areaCircle(x float64) {
	area := math.Pi * x * x
	fmt.Println("Area of the circle = ", area)
}

func areaSquare(x float64) {
	area := x * x
	fmt.Println("Area of the square = ", area)
}
