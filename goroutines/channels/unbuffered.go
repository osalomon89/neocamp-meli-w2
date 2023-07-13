// package main

// import (
// 	"fmt"
// 	"time"
// )

// // channels: los canales son una estructura de datos que proporciona una forma de comunicación y sincronización entre goroutines.
// // Los canales permiten el envío y la recepción de valores entre goroutines, asegurando una comunicación segura y coordinada.
// // https://go.dev/tour/concurrency/2

// // Canal sin buffer (unbuffered channel): El canal sin buffer bloqueará tanto la goroutine emisora como la receptora hasta que se realice la comunicación, asegurando así una sincronización coordinada.
// // Canal con buffer (buffered channel): el canal con buffer permite el envío de valores sin bloquear la goroutine emisora siempre y cuando haya espacio disponible en el buffer.

// func main() {
// 	start := time.Now()
// 	c := make(chan string)

// 	for i := 0; i < 5; i++ {
// 		go getOrder(i, c)
// 	}

// 	for i := 0; i < 5; i++ {
// 		msg := <-c
// 		fmt.Print(msg)
// 	}

// 	close(c)

// 	fmt.Println("Finished")
// 	fmt.Println(time.Since(start))
// }

// func getOrder(id int, c chan string) {
// 	time.Sleep(1 * time.Second)
// 	msg := fmt.Sprintf("The API call, number %d, was executed\n", id)

// 	c <- msg
// }
