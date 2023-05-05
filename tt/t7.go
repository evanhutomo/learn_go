package main

import "fmt"

func main() {
	// common switch
	for i := 0; i < 3; i++ {
		switch {
		case i == 0:
			fmt.Println("zero")
		case i == 1:
			fmt.Println("num: ", i)
		default:
			break
		}
	}

	// blank switch
	words := []string{"hi", "salutations", "hello"}
	for _, words := range words {
		switch wordLen := len(words); {
		case wordLen < 5:
			fmt.Println("less than 5")
		case wordLen > 10:
			fmt.Println("more than 10")
		}
	}
}
