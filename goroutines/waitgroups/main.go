package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go getOrder(i, &wg)
	}

	wg.Wait()
	fmt.Println("Finished")
	fmt.Println(time.Since(start))
}

func getOrder(id int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Printf("The API call, number %d, was executed\n", id)

	wg.Done()
}
