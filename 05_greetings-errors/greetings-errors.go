package greetings

import (
	"errors"
	"fmt"
)

// Greet returns a greeting for the named person.
func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf("Hi %v!", name)

	return message, nil
}
