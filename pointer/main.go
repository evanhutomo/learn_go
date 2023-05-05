package main

import "fmt"

func changeString(s *string) {
	fmt.Println(s)
	var newValue string
	newValue = "hutomo"
	*s = newValue
}

func main() {
	var test string
	test = "evan"
	fmt.Println(test)
	changeString(&test)
	fmt.Println(test)
}
