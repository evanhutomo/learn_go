package main

import "fmt"

type mailer interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("sending email")
}

type admin struct {
	user
	level int
}

func main() {
	user := user{
		name:  "evan hutomo",
		email: "evanhutomo@gmail.com",
	}

	ad := &admin{
		user:  user,
		level: 8,
	}

	ad.notify()
}
