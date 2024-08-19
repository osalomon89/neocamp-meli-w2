package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// El statement select se utiliza ampliamente para manejar múltiples canales y
// coordinar operaciones de lectura y escritura en tiempo real.
// Es útil cuando se trabaja con comunicación y sincronización entre múltiples goroutines,
// y permite tomar acciones basadas en qué canal está listo para la operación.

func main() {
	c := make(chan string)
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doHttpRequest(ctx, i, c, &wg)
	}

loop:
	for {
		select {
		case msg := <-c:
			fmt.Print(msg)
		case <-ctx.Done():
			fmt.Println("Timeout exceeded, stopping all operations.")
			break loop
		}
	}

	wg.Wait()
	fmt.Println("All calls were executed")
	close(c)
}

func doHttpRequest(ctx context.Context, id int, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(2500)
	time.Sleep(time.Duration(n) * time.Millisecond)

	msg := fmt.Sprintf("The API call, number %d, was executed\n", id)

	select {
	case c <- msg:
		// Si el contexto no ha expirado, enviamos el mensaje
	case <-ctx.Done():
		// Si el contexto ha expirado, simplemente salimos
		fmt.Printf("Timeout: #%d\n", id)
		return
	}
}
