package main

import (
	"fmt"
	"log"

	greetings "crash-course_01/05_greetings-errors"
)

func main() {
	message, err := greetings.Greet("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
