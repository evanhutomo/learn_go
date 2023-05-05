package greetings

import "fmt"

// Hello returns a greeting for the named 
func Hello(name string) string {
	// return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
