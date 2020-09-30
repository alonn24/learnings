package main

import (
	"greetings"
	"log"
)

func hello() {
	// Lambda
	getNames := func() []string {
		return []string{"Gladys", "Samantha", "Darrin"}
	}
	// slice of names
	names := getNames()
	messages, err := greetings.Hellos(names...)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("> hello: %v\n", messages)
}
