package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// hello returns a greeting for the named person.

func Hello(name string) string {
	if name == "" {
		return name, errors.New("Nope not today sir!")
	}
	// Return a greeting that embeds the name in a message, using randomFormat
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
// init sets initial values for variables used in the function.
func init() {
		rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats.
	formats := [] string{
		"%v erm",
		"dafuq %v",
		"%v %v %v",
	}
	// return randomly selected message formats
	return formats[rand.Intn(len(formats))]
}

func Hello2(name2 string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Whazzzzzaaaa!", name2)
	return message
}

func Hello3(name2 string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Goodbye", name2)
	return message
}
