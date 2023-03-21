package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "api-key", "my-super-secret-api-key")
}

func doSomethingCool(ctx context.Context) {
	apiKey := ctx.Value("api-key")
	fmt.Println("ctx.Value: ", apiKey)
}

func main332() {
	fmt.Println("--Go Context Example--")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomethingCool(ctx)
}
