package main

import "fmt"

func modificarArray(array [3]int) {
	array[0] = 100
	fmt.Println("array em func: ", array)
}

func modificarSlice(slice []int) {
	slice[0] = 100
}

func main() {
	// Array
	array := [3]int{1, 2, 3}
	fmt.Println("Array antes:", array)
	modificarArray(array)
	fmt.Println("Array depois:", array)

	// Slice
	slice := []int{1, 2, 3, 4}
	fmt.Println("Slice antes:", slice)
	modificarSlice(slice)
	fmt.Println("Slice depois:", slice)
}
