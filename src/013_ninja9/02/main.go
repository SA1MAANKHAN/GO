package main

import "fmt"

type buggati struct {
	speed int
}
type car interface {
	seats(x int)
	//  {
	//   if x < 4{
	//     fmt.Println("it's two seater")
	//   }else{
	//     fmt.Println("it's more than two seats")
	//   }
	// }
}

func (b *buggati) seats() string {
	// fmt.Println(" i am num of seats")\
	m := "i am a seat"
	return m
}
func info(c car) {
	fmt.Println("i am a car and ", c.seats(""))
}

func main() {
	c := buggati{300}
	info(c)
}
