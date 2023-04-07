package ex_context

import (
	"fmt"
)

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}
