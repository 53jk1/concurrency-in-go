package ex_context

import (
	"context"
	"fmt"
)

func printGreetingWithContext(ctx context.Context) error {
	greeting, err := genGreetingWithContext(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}
