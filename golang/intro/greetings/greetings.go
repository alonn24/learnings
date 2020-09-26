package greetings

import (
	"errors"
	"fmt"
)

// Hello - Function with capital letter is exported
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// Hellos - do something
func Hellos(names ...string) (map[string]string, error) {
	messages := make(map[string]string)
	// loop over slice of names
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
