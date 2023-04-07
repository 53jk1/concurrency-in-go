package ex_context

import (
	"context"
	"fmt"
)

func printFarewellWithContext(ctx context.Context) error {
	farewell, err := genFarewellWithContext(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}
