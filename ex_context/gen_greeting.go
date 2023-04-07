package ex_context

import (
	"fmt"
)

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "Hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}
