package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Doing an HTTP request...")

	result := make(chan string)
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	start := time.Now()

	// Agregamos el número de goroutines al WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go DoHttpRequest5(ctx, result, &wg, i)
	}

	// Escuchamos el resultado de la primera goroutine que termine
	msg := <-result
	elapsed := time.Since(start)

	log.Printf("The HTTP request took %s", elapsed)
	log.Printf("The HTTP response was: %s", msg)

	// Esperamos a que todas las goroutines finalicen antes de salir
	wg.Wait()
	close(result)
}

// DoHttpRequest performs an HTTP request that can take between 0 and 500ms to be done
func DoHttpRequest5(ctx context.Context, result chan<- string, wg *sync.WaitGroup, i int) {
	defer wg.Done() // Decrementa el WaitGroup counter al terminar la goroutine

	// Simulamos una solicitud HTTP
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
