package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	var message string
	message = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you %v",
		"Wasssup, bitch %v",
		"Fuck around and find out.",
	}

	return formats[rand.Intn(len(formats))]

}
