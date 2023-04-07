package ex_context

import (
	"context"
	"time"
	"fmt"
)

func genGreetingWithContext(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	
	switch locale, err := localeWithContext(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "Hello", nil
		
	}
	return "", fmt.Errorf("unsupported locale")
}
