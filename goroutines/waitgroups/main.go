package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	numbers := []int{1, 2, 3, 4, 5}

	wg := sync.WaitGroup{}
	wg.Add(len(numbers))

	for _, number := range numbers {
		go getOrder(number, &wg)
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
