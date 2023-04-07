package ex_context

import (
	"sync"
	"context"
	"fmt"
)

func ConcurrentPrintGreetAndFarewellWithContext() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreetingWithContext(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewellWithContext(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
			cancel()
		}
	}()
	
	wg.Wait()
}
