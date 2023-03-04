package main

import (
	"fmt"
	"log"

	"geekslayr.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Azeem", "Faizan", "Habib"}

	// var message string
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	fmt.Println(quote.Glass())
}
