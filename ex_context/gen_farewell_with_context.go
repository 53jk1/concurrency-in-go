package ex_context

import (
	"context"
	"fmt"
)

func genFarewellWithContext(ctx context.Context) (string, error) {
	switch locale, err := localeWithContext(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "Goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}
