package greetings

import (
	"regexp"
	"testing"
)

func TestGreetName(t *testing.T) {
	name := "John Doe"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	message, err := Greet("")

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
