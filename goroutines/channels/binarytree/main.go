package main

// import (
// 	"fmt"

// 	"golang.org/x/tour/tree"
// )

// // Walk recorre el árbol t enviando todos los valores
// // del árbol al canal ch.
// func Walk(t *tree.Tree, ch chan int) {
// 	// Función auxiliar para recorrer el árbol en orden
// 	var walker func(t *tree.Tree)
// 	walker = func(t *tree.Tree) {
// 		if t == nil {
// 			return
// 		}
// 		walker(t.Left)
// 		ch <- t.Value
// 		walker(t.Right)
// 	}

// 	walker(t)
// 	close(ch)
// }

// // Same determina si los árboles t1 y t2 contienen los mismos valores.
// func Same(t1, t2 *tree.Tree) bool {
// 	ch1, ch2 := make(chan int), make(chan int)
// 	go Walk(t1, ch1)
// 	go Walk(t2, ch2)

// 	for {
// 		v1, ok1 := <-ch1
// 		v2, ok2 := <-ch2

// 		if ok1 != ok2 || v1 != v2 {
// 			return false
// 		}
// 		if !ok1 {
// 			return true
// 		}
// 	}
// }

// func main() {
// 	// Prueba de Walk
// 	ch := make(chan int)
// 	go Walk(tree.New(1), ch)
// 	fmt.Println("Valores del árbol:")
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%d ", <-ch)
// 	}
// 	fmt.Println()

// 	// Prueba de Same
// 	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(tree.New(1), tree.New(1)))
// 	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))
// }
