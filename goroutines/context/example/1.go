package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main1() {
	fmt.Println("Doing an HTTP request...")

	result := make(chan string)

	start := time.Now()
	go DoHttpRequest1(result, 1)
	go DoHttpRequest1(result, 2)
	go DoHttpRequest1(result, 3)
	go DoHttpRequest1(result, 4)
	go DoHttpRequest1(result, 5)

	msg := <-result
	elapsed := time.Since(start)

	log.Printf("The HTTP request took %s", elapsed)
	log.Printf("The HTTP response was: %s", msg)

	time.Sleep(2 * time.Second)
}

// DoHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoHttpRequest1(result chan<- string, i int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(500)
	time.Sleep(time.Duration(n) * time.Millisecond)

	result <- fmt.Sprintf("Goroutine finished #%d\n", i)
}
