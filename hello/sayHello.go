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
	// var message string
	message, err := greetings.Hello("Azeem")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(quote.Glass())
}
