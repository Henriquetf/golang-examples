package main

import (
	"fmt"
	"log"

	"octopus/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Rocky", "Grace", "Astrophage"}

	message, err := greetings.MultiGreet(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
