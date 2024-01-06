package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	messageFormats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail, %v! Well met",
	}
	lenOfMessageFormats := len(messageFormats)
	elementIndexSelector := rand.Intn(lenOfMessageFormats)
	res := fmt.Sprintf("lenOfMessageFormats is %v, elementIndexSelector is %v", lenOfMessageFormats, elementIndexSelector)
	fmt.Println(res)

	return messageFormats[elementIndexSelector]
}

// Hello returns a greeting for the named person.
func Hello(name string) (message string, err error) {
	if name == "" {
		message = ""
		err = errors.New("empty name")
		return
	}

	message = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	//nameMessageKeyValuePairs := make(map[string]string) // Can we create maps any other way?
	nameMessageKeyValuePairs := map[string]string{} // Yip
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		nameMessageKeyValuePairs[name] = message
	}
	return nameMessageKeyValuePairs, nil
}
