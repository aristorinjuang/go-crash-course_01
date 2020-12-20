package greetings

import "fmt"

// Greet returns a greeting for the named person.
func Greet(name string) string {
	message := fmt.Sprintf("Hi %v!", name)

	return message
}
