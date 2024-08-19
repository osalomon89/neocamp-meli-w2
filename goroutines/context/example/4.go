package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main4() {
	fmt.Println("Doing an HTTP request...")

	result := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	start := time.Now()
	go DoHttpRequest4(ctx, result, 1)
	go DoHttpRequest4(ctx, result, 2)
	go DoHttpRequest4(ctx, result, 3)
	go DoHttpRequest4(ctx, result, 4)
	go DoHttpRequest4(ctx, result, 5)

	msg := <-result
	cancel()
	elapsed := time.Since(start)

	log.Printf("The HTTP request took %s", elapsed)
	log.Printf("The HTTP response was: %s", msg)

	time.Sleep(2 * time.Second)
}

// DoHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoHttpRequest4(ctx context.Context, result chan<- string, i int) {
	// Do an HTTP request synchronously
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(500)
	time.Sleep(time.Duration(n) * time.Millisecond)

	response := fmt.Sprintf("Goroutine finished #%d\n", i)

	select {
	case result <- response:
		fmt.Printf("Goroutine finished #%d\n", i)
	case <-ctx.Done():
		fmt.Printf("Timeout: #%d\n", i)
	}
}
