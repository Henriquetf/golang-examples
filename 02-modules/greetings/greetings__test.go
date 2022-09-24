package greetings

import (
	"regexp"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "CatDog"
	message, err := Greet(name)
	want := regexp.MustCompile(`\b` + name + `\b`)
	matches := want.MatchString(message)

	if !matches || err != nil {
		t.Fatalf(`Greet(%q) = %q, %v, want match for %#q, nil`, name, message, err, want)
	}
}

func TestGreetCannotBeEmpty(t *testing.T) {
	name := ""
	message, err := Greet(name)
	want := ""

	if message != want || err == nil {
		t.Fatalf(`Hello(%q) = %q, %v; want %q, error`, name, message, err, want)
	}
}

func TestMultiGreet(t *testing.T) {
	names := []string{"What", "The", "Dog", "Doin"}

	messages, err := MultiGreet(names)

	if messages == nil || err != nil {
		t.Fatalf(`MultiGreet(names) = %q, %v; want messages, nil`, messages, err)
	}
}

func TestMultiGreetCannotHaveEmptyNames(t *testing.T) {
	names := []string{"What", "The", "", "Dog", "Doin"}

	messages, err := MultiGreet(names)

	if messages != nil || err == nil {
		t.Fatalf(`MultiGreet(names) = %q, %v; want nil, error`, messages, err)
	}
}
