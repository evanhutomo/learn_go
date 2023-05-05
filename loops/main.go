package main

import "log"

type User struct {
	Name string
	Age  int
}

func main() {
	var data []User
	data = append(data, User{"Evan", 32})
	data = append(data, User{"Ekwi", 36})
	data = append(data, User{"Edwin", 30})

	for _, d := range data {
		log.Println(d.Name, d.Age)
	}
}
