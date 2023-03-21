package main

import (
	"context"
)

func doSomething(ctx context.Context) error {
	for {
		// do something
		select {
		case <-ctx.Done(): // closes when the caller cancels the ctx
			return ctx.Err() // has a value on context cancellation
		}
	}
}

func mainWithCancel2() {
	ctx := context.Background()
	ctx, stop := context.WithCancel(ctx)
	// OR ctx, stop := context.WithTimeout(ctx, time.Second*3)
	// OR ctx, stop := context.WithDeadline(ctx, time.Now().Add(time.Second*5))

	// use the context to do some work in the background
	// other operation that results in a error
	if doSomething(ctx) != nil {
		stop() // send the cancellation signal to all functions using ctx
	}
}
