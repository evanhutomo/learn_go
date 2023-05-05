package main

import "log"

type PeopleAddress struct {
	Address string
	Code    int
}

func TestMap() {
	myAddress := make(map[string]PeopleAddress)

	data := PeopleAddress{
		Address: "Jl. Kepuh IX/8",
		Code:    65148,
	}

	myAddress["saya"] = data
	log.Println(myAddress["saya"].Address)
}

func main() {
	TestMap()
}
