package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Initilazing properties of logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Initializing slice of names
	names := []string{"Alex", "Anca", "Claudiu"}

	// Get a greeting message and handle for errors
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no errors returned, print message to console
	fmt.Println(message)
}
