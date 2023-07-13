package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel
	baton := make(chan int)
	quit := make(chan bool)

	// First runner to his mark
	go Runner(baton, quit)

	// Start the race
	baton <- 1

	// Give the runners time to race
	<-quit
}

func Runner(baton chan int, quit chan bool) {
	var newRunner int

	// Wait to receive the baton
	runner := <-baton

	// Start running around the track
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// New runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton, quit)
	}

	// Running around the track
	time.Sleep(50 * time.Millisecond)

	// Is the race over
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		quit <- true
		return
	}

	// Exchange the baton for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}
