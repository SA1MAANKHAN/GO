package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Userone",
		Age:   20,
	}
	u2 := user{
		First: "Usertwo",
		Age:   21,
	}
	u3 := user{
		First: "Userthree",
		Age:   22,
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)

	jsonUser, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonUser))

}
