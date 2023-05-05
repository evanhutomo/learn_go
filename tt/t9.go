package main

import (
	"fmt"
	"sort"
)

func main() {
	type GroupStudent struct {
		Name string
		Age  int
	}

	people := []GroupStudent{
		{"evan", 32},
		{"ekwi", 36},
		{"agung", 30},
	}
	fmt.Println(people)

	// sort by name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println(people)

	// sort byage
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
