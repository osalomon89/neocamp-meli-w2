package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	var inputNumber int

	for {
		fmt.Printf("\ningrese el numero: ")
		fmt.Scan(&inputNumber)

		if randomNumber > inputNumber {
			fmt.Println("El número es menor")
			continue
		}

		if randomNumber < inputNumber {
			fmt.Println("El número es mayor")
			continue
		}

		fmt.Printf("\n*****Adivinaste: %d******\n", randomNumber)
		break
	}
}
