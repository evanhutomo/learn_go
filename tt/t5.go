package main

import "fmt"

func main() {
	// test for loop
	// 1. A complete, c style for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 2. A condition-only for
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
	// 3. an infinite for
	for {
		// things to do in the loop
		if !CONDITION {
			break
		}
	}
	// 4. for-range
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	// ignoring the key in a for-range loop
	evenVals := []int{2, 4, 6}
	for _, v := range evenVals {
		fmt.Println(v)
	}
}
