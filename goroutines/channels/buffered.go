package main

// func main() {
// 	start := time.Now()
// 	c := make(chan string, 5)

// 	for i := 0; i < 5; i++ {
// 		go getOrder(i, c)
// 	}

// 	msg1 := <-c
// 	msg2 := <-c
// 	msg3 := <-c
// 	msg4 := <-c
// 	msg5 := <-c

// 	fmt.Println(msg1, msg2, msg3, msg4, msg5)

// 	close(c)

// 	fmt.Println("Finished")
// 	fmt.Println(time.Since(start))
// }

// func getOrder(id int, c chan string) {
// 	time.Sleep(1 * time.Second)
// 	msg := fmt.Sprintf("The API call, number %d, was executed\n", id)

// 	c <- msg
// }
