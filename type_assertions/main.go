package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	// panic if type assertion is wrong
	//i2 := i.(string)
	fmt.Println(i2 + 1)

	m := map[string]int {
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
}
