package main

import "fmt"

func main() {
	var arr [6]int 
	capArr := cap(arr)
	fmt.Println("the capacity of the array is: ")
	fmt.Println(capArr)

	slice := make([]string, 0, 10)
	capSlice := cap(slice)
	fmt.Println("The capacity of the slice is:")
	fmt.Println(capSlice)

	channel := make(chan string, 5)
	capChannel := cap(channel)
	fmt.Println("The capacity of the slice is:")
	fmt.Println(capChannel)

	var pointer *[20] string
	capPointer := cap(pointer)
	fmt.Println("The capacity of the pointer is: ")
	fmt.Println(capPointer)

	m := map[string]int {
		"hello": 5,
		"world": 0,
	}
	v, o := m["hello"]
	fmt.Println(v, o)

	v, o = m["world"]
	fmt.Println(v, o)

	v, o = m["goodnbye"]
	fmt.Println(v, o)

}
