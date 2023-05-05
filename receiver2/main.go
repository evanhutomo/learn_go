package main

import "log"

type User struct {
	FirstName string
	Age       int
}

func (u *User) printUserData() (string, int) {
	if u.FirstName == "" {
		u.FirstName = "DefaultName"
	}
	return u.FirstName, (u.Age - 3)
}

func main() {
	var myUser User
	myUser.FirstName = "Evan"
	myUser.Age = 32
	myUser2 := User{
		Age: 32,
	}

	log.Println(myUser.printUserData())
	log.Println(myUser2.printUserData())
}
