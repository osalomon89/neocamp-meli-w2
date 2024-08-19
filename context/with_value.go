package main

import (
	"context"
	"fmt"
)

func doSomethingCool(ctx context.Context) {
	apiKey := ctx.Value("api-key")
	name := ctx.Value("nome")

	fmt.Println("ctx.Value: ", apiKey)
	fmt.Println("name: ", name)
}

func main() {
	fmt.Println("--Go Context Example--")
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "api-key", "my-super-secret-api-key")
	ctx = context.WithValue(ctx, "nome", "oscar")

	doSomethingCool(ctx)
}
