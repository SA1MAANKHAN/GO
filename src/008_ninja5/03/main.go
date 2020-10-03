package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{vehicle: vehicle{doors: 4, color: "red"},
		fourwheel: true,
	}
	s := sedan{vehicle: vehicle{doors: 2, color: "green"},
		luxury: true,
	}
	fmt.Println(t, s)

}
