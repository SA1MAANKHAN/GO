package main

import "fmt"

type Person struct {
	first_name string
	last_name  string
	fav_flavs  []string
}

func main() {

	p1 := Person{
		first_name: "Salmaan",
		last_name:  "Khan",
		fav_flavs:  []string{"belgian chocochip", "straberry", "choco"},
	}
	p2 := Person{
		first_name: "Leo",
		last_name:  "Nardo",
		fav_flavs:  []string{"chocochip", "dry fruit", "vanilla"},
	}
	m := map[string]Person{p1.last_name: p1, p2.last_name: p2}

	fmt.Println("this is map \n", m)

	for k, v := range m {
		fmt.Println(k, v)
		for i, val := range v.fav_flavs {
			fmt.Println(i, val)
		}
	}

}
