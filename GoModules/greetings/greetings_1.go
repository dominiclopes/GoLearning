package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Inside the init of greetings_1.go")
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (message string, err error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message = fmt.Sprintf(randFormat(), name)
	return
}

func Hellos(names []string) (messages map[string]string, err error) {
	messages = make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return map[string]string{}, err
		}
		messages[name] = message
	}
	return
}

func randFormat() string {
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you %v.",
		"Hail %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
