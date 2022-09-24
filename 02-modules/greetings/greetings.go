package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func MultiGreet(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greet(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func init() {
	// 1 nano = 1e-9 second
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello, %s :)",
		"Great to unsee you, %s!",
		"Hello, %s! I'm in need of some Astrophage",
	}

	formatIndex := rand.Intn(len(formats))

	return formats[formatIndex]
}
