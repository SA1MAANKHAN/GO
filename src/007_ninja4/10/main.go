package main

import "fmt"

func main() {
	m := map[string][]string{"Sharma": {"football", "baseball", "Badminton"}, "Verma": {"table tennis", "chess", "music"}}
	m["Khan"] = []string{"football", "meditaion", "books"}

	delete(m, "Sharma")
	fmt.Print(m)
}
