package main

import (
	"fmt"
	"log"

	greetings "crash-course_01/07_greetings-test"
)

func main() {
	message, err := greetings.Greet("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
