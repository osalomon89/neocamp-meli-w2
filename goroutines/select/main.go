package main

import (
	"fmt"
	"time"
)

func worker1(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "Mensaje del worker 1"
}

func worker2(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "Mensaje del worker 2"
}

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go worker1(canal1)
	go worker2(canal2)

	// Usamos select para esperar los mensajes de los workers
	select {
	case mensaje1 := <-canal1:
		fmt.Println(mensaje1)
	case mensaje2 := <-canal2:
		fmt.Println(mensaje2)
	}

	fmt.Println("All calls were executed")
	close(canal1)
	close(canal2)
}
