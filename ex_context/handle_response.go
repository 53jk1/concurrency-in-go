package ex_context

import (
	"context"
	"fmt"
)

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (%v)\n",
		ctx.Value("userID"),
		ctx.Value("authToken"),
	)
}
