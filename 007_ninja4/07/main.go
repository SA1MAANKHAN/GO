package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken ,not stirred"}, {"Miss", "MonneyPenny", "Heloooo ,James"}}

	fmt.Println("ranging over records")
	for i := 0; i < 2; i++ {
		fmt.Println(x[i])
	}
	fmt.Println("ranging over data")
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Println(x[i][j])
		}
	}
}
