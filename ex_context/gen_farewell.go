package ex_context

import (
	"fmt"
)

func genFarewell(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "Goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}
