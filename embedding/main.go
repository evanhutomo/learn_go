package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // embedded field
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return nil
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	// var eFail Employee = m // compilation error
	var eOK Employee = m.Employee // compilation ok
	fmt.Println(eOK)  // prints 12345
	fmt.Println(m.ID)  // prints 12345
	fmt.Println(m.Description())  // prints Bob Bobson (12345)
}
