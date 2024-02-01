package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	testContext()
}

func testContext() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	go sleepingFunc(ctx)
	time.Sleep(500 * time.Millisecond)
	if _, ok := <-ctx.Done(); !ok {
		fmt.Println("Context channel was closed because of timeout")
	}
	fmt.Println("Finished execution")
}

func sleepingFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Finished execution after context done")
			return
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Something")
		}
	}
}
