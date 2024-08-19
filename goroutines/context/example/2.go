package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main2() {
	fmt.Println("Doing an HTTP request...")

	result := make(chan string)

	start := time.Now()
	go DoHttpRequest2(result, 1)
	go DoHttpRequest2(result, 2)
	go DoHttpRequest2(result, 3)
	go DoHttpRequest2(result, 4)
	go DoHttpRequest2(result, 5)

	msg := <-result
	elapsed := time.Since(start)

	log.Printf("The HTTP request took %s", elapsed)
	log.Printf("The HTTP response was: %s", msg)

	time.Sleep(2 * time.Second)
}

// DoHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoHttpRequest2(result chan<- string, i int) {
	// Do an HTTP request synchronously
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(500)
	time.Sleep(time.Duration(n) * time.Millisecond)

	response := fmt.Sprintf("Goroutine finished #%d\n", i)

	select {
	case result <- response:
		fmt.Printf("Goroutine finished #%d\n", i)
	case <-time.After(1 * time.Second):
		fmt.Printf("Timeout #%d\n", i)
	}
}
