package main

import "fmt"

var favsport string

func main() {
	favsport = "football"
	switch favsport {
	case "tennis":
		fmt.Print("know Roger Fedrer?")
	case "football":
		fmt.Print("know Messi?")
	case "cricket":
		fmt.Print("know Tendulkar")
	}

}
