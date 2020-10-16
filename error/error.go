package error

import (
  "errors"
  "fmt"
)

// Hello returns a greeting for the named person.

func Hello(name string) (string, error){
  // If no name was given, return an error with a message.
  if name == ""{
    return "", errors.New("dafuq name was empty man?")
  }
  // If a name was received, return a value that embeds the name
  // In a greeting message.
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message, nil
}
