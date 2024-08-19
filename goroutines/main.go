package main

import (
	"fmt"
	"time"
)

// Paralelismo: se refiere a la capacidad de realizar varias tareas simultáneamente.
// En un contexto de programación, implica dividir una tarea grande en subtareas más pequeñas y
// ejecutarlas al mismo tiempo utilizando múltiples recursos de procesamiento.
// El objetivo del paralelismo es mejorar la eficiencia y reducir el tiempo de ejecución total de un programa.
// Puede lograrse mediante el uso de múltiples hilos, procesadores o incluso sistemas distribuidos.

// Concurrencia: La concurrencia se refiere a la capacidad de administrar múltiples tareas que se ejecutan de manera independiente y aparentemente simultánea.
// A diferencia del paralelismo, la concurrencia no implica que las tareas se ejecuten al mismo tiempo exactamente,
// sino que pueden progresar en solapamiento o de forma intercalada.
// En lugar de dividir una tarea en subprocesos, la concurrencia permite que varias tareas se ejecuten de forma independiente, aunque puedan compartir recursos.

func main() {
	start := time.Now()

	for i := 0; i < 5; i++ {
		//getOrder(i)
		go getOrder(i)
	}

	fmt.Println("FOR finished")
	time.Sleep(2 * time.Second)

	fmt.Println("Finished")
	fmt.Println(time.Since(start))
}

func getOrder(id int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Task number %d, was executed\n", id)
}
