package main

import (
	"fmt"
	"geekslayr.com/greetings"
	"log"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// var message string
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(quote.Glass())
}
