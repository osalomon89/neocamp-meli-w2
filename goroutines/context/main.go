package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Doing an HTTP request...")

	start := time.Now()
	result := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	go doHttpRequest(ctx, result, 1)
	go doHttpRequest(ctx, result, 2)
	go doHttpRequest(ctx, result, 3)
	go doHttpRequest(ctx, result, 4)
	go doHttpRequest(ctx, result, 5)

	response := <-result
	cancel()

	elapsed := time.Since(start)

	log.Printf("The HTTP request took %s", elapsed)
	log.Printf("The HTTP response was: %s", response)

	time.Sleep(3 * time.Second)
}

// doHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func doHttpRequest(ctx context.Context, result chan<- string, i int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(500)
	time.Sleep(time.Duration(n) * time.Millisecond)

	response := fmt.Sprintf("Goroutine finished #%d\n", i)

	select {
	case result <- response:
		fmt.Printf("Goroutine finished #%d\n", i)
	case <-ctx.Done():
		fmt.Printf("Goroutine finished #%d\n", i)
	}
}
