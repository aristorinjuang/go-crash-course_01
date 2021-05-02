package main

import (
	"fmt"

	greetings "crash-course_01/03_greetings"
)

func main() {
	message := greetings.Greet("John Doe")

	fmt.Println(message)
}
