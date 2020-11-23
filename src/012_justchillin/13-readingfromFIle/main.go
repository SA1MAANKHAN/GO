package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("hello")
	data, err := ioutil.ReadFile("/home/salmaan/Documents/GitHub/GO/src/012_justchillin/13-readingfromFIle/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
