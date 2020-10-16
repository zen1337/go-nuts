package mgreet

import (
  "errors"
  "fmt"
  "math/rand"
  "time"
)

// Hellos return a map that associates each of the named people
// with a greeting message.

func Hello(name string) (string, error) {
  if name == "" {
    return name, errors.New("dafuq???? %v do you see here smn bicz?", name)
  }
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

func Hellos(name []string) (map[string]string, error) {
  // A map to associate names with messages.
  messages := make(map[string]string)
  // Loop through the receives slice of names, calling
  // the hello function the get a message for each name.
  for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }
    // In the map, associate the retrived message with
    // the name.
    messages[name] = message
  }
  return message, nil
}

// Init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}
