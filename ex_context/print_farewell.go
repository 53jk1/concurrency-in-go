package ex_context

import (
	"fmt"
)

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}
