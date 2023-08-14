package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.WithValue(context.Background(), 1, "value_in_context"),
		6*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("context has ended,", ctx.Err().Error())
	case <-time.After(5 * time.Second):
		fmt.Println("Executing after 5 seconds with context value,", ctx.Value(1))
	}
}
