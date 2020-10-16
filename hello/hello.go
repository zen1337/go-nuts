package main

import (
	"fmt"
	"log"
	"n3.com/greetings"
	"n3.com/error"
	"n3.com/mgreet"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("error.go: ")
	log.SetFlags(log.LstdFlags)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	message := greetings.Hello("Adam")
	fmt.Println(message)
	message2 := greetings.Hello2("Adam")
	fmt.Println(message2)
	message3 := greetings.Hello3("Adam")
	fmt.Println(message3)
	// Request a greetings message.
	message, err := error.Hello("")
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}
	// Get a greeting message and print it.
//	message, err := error.Hello("")
//	fmt.Println(message)
//	message := greetings.Hello2("Ania")
	fmt.Println(message)
//	message3 := greetings.Hello3("Ania")
//	fmt.Println(message3)
}
