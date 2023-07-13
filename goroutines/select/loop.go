// package main

// import (
// 	"fmt"
// 	"time"
// )

// // El statement select se utiliza ampliamente para manejar múltiples canales y
// // coordinar operaciones de lectura y escritura en tiempo real.
// // Es útil cuando se trabaja con comunicación y sincronización entre múltiples goroutines,
// // y permite tomar acciones basadas en qué canal está listo para la operación.

// func main() {
// 	c := make(chan string)
// 	quit := make(chan bool)

// 	var counter int
// 	for i := 0; i < 5; i++ {
// 		go getOrder(i, c, quit, &counter)
// 	}

// loop:
// 	for {
// 		select {
// 		case msg := <-c:
// 			fmt.Print(msg)
// 		case <-quit:
// 			fmt.Println("You've exhaust my patience")
// 			break loop
// 		}
// 	}

// 	fmt.Println("All calls were executed")
// 	close(c)
// 	close(quit)
// }

// func getOrder(id int, c chan string, quit chan bool, counter *int) {
// 	time.Sleep(1 * time.Second)

// 	msg := fmt.Sprintf("The API call, number %d, was executed\n", id)
// 	c <- msg

// 	*counter++
// 	if *counter >= 5 {
// 		quit <- true
// 	}
// }
