package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Note: For functions to be exported they must be UPPERCASE
//		 lowercase functions are private to the package

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name give, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// map -> map[key-type]value-type
	messages := make(map[string]string)

	// "_" is the blank identifier and it is used as a ignored variable
	// range returns (index, item) and index is not being used to we use "_"
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// Assocate retrived message with name
		messages[name] = message
	}

	return messages, nil

}

// Creates random format string for greetings
func randomFormat() string {
	// Defining "slice" (i.e. cpp vector) of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
