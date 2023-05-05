// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// // example of method
// func (p Person) String() string {
// 	return fmt.Sprintf("%s %d", p.name, p.age)
// }

// func main() {
// 	p := Person{
// 		name: "evan",
// 		age:  32,
// 	}
// 	output := p.String()
// 	fmt.Println(output)
// }

package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}
