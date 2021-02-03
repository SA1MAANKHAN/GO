package main

import "fmt"

var x []string

func main() {
	x = []string{"Andhra Pradesh", "Arunachal Pradesh", "Assam", "Bihar", "Chhattisgarh", "Goa", "Gujarat", "Haryana", "Himachal Pradesh", "Jharkhand", " Karnataka ", "Kerala", "Madhya Pradesh", "Maharashtra", "Manipur", "Meghalaya", "Mizoram", "Nagaland ", "Odisha", "Punjab", "Rajasthan", "Sikkim", "Tamil Nadu", "Telangana ", "Tripura", "Uttar Pradesh ", "Uttarakhand ", "West Bengal"}
	// fmt.Print(len(x))

	fmt.Println("Length of the silce is", len(x))

	for i := 1; i <= len(x); i++ {
		fmt.Printf("%v %v\n", i, x[i-1])

	}
}
