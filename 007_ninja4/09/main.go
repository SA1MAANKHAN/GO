package main

import "fmt"

func main() {
	m := map[string][]string{"kumar": {"football", "baseball", "Badminton"}, "Verma": {"table tennis", "chess", "music"}}
	m["Khan"] = []string{"football", "meditaion", "books"}
	fmt.Print(m)

}
