package main

import (
	"fmt"
)

type Creature struct {
	Species string
}

func main() {
	creature := Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}

/*func main() {
	w := wallet.Wallet{}
	w.Deposit(10)
	w.Deposit(50)

	fmt.Printf("Balance: %d\n", w.Balance())
}*/
