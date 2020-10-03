package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	val, err := sqrt(-10.23)
	if err != nil {
		// log.Println(err)
		log.Fatalln(err)
		return
	}
	fmt.Println("here is your answer : ", val)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Error ! , you entered a negative number --- value was : %v", f)
	}
	return 42, nil
}
